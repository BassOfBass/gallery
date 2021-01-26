package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

// CreateThumbnail creates a thumbnail image if it doesn't already exist
func CreateThumbnail(fpath, tpath string, overwrite bool) error {
	_, err := os.Stat(fpath)
	if err != nil {
		return err
	}

	_, err = os.Stat(tpath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if err == nil && !overwrite {
		return nil
	}

	f, err := os.Open(fpath)
	if err != nil {
		return err
	}
	defer f.Close()

	img, format, err := image.Decode(f)
	if err != nil {
		return err
	}

	t := resize.Thumbnail(ThumbnailDimensions.Width, ThumbnailDimensions.Height, img, resize.Bilinear)

	out, err := os.Create(tpath)
	if err != nil {
		return err
	}
	defer out.Close()

	switch format {
	case "gif":
		err = gif.Encode(out, t, nil)
	case "jpeg":
		err = jpeg.Encode(out, t, &jpeg.Options{Quality: 10})
	case "png":
		err = png.Encode(out, t)
	default:
		err = fmt.Errorf("unknown format: %s", format)
	}
	if err != nil {
		return fmt.Errorf("error encoding thumbnail: %v", err)
	}

	return nil
}
