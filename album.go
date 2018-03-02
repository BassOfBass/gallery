package main

import (
	"crypto/md5"
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Album represents an directory containing images to be served
type Album struct {
	Name   string
	Images Images
}

// NewAlbum returns a new album for a given directory
func NewAlbum(d string) (Album, error) {
	a := Album{}
	ds, err := os.Stat(d)
	if err != nil {
		return a, fmt.Errorf("cannot stat directory %s: %v", d, err)
	}
	a.Name = ds.Name()
	a.Images = Images{}

	filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Ignore if the file is already a thumbnail
		if strings.Contains(info.Name(), "_thumb") {
			return nil
		}
		// If the file is an allowed file extension then include it
		for _, e := range AllowedExtensions {
			if strings.HasSuffix(strings.ToLower(info.Name()), e) {
				thumb := fmt.Sprintf("%s_thumb%s", strings.TrimSuffix(path, e), e)
				f, err := os.Open(path)
				if err != nil {
					return fmt.Errorf("error opening image: %v", err)
				}
				ic, _, err := image.DecodeConfig(f)
				if err != nil {
					return fmt.Errorf("error decoding image config: %v", err)
				}
				id := Image{
					Path:          path,
					ThumbnailPath: thumb,
					Md5:           fmt.Sprintf("%x", md5.Sum([]byte(path))),
					ThumbnailMd5:  fmt.Sprintf("%x", md5.Sum([]byte(thumb))),
					LastModified:  info.ModTime(),
					Width:         ic.Width,
					Height:        ic.Height,
				}
				a.Images = append(a.Images, id)
				break
			}
		}
		return nil
	})

	// Sorting images by modfied time, newest first
	sort.Sort(sort.Reverse(a.Images))

	genThumbnails(a)

	return a, nil
}

// genThumbnails generates missing thumbnails
func genThumbnails(a Album) error {
	log.Println("Generating missing thumbnails...")
	for _, img := range a.Images {
		err := CreateThumbnail(img.Path, img.ThumbnailPath, false)
		if err != nil {
			return fmt.Errorf("error creating thumbnail: %v", err)
		}
	}
	return nil
}
