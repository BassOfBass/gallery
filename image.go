package main

import "time"

// Image represents a single album image
type Image struct {
	Path          string
	ThumbnailPath string
	Md5           string
	ThumbnailMd5  string
	LastModified  time.Time
	Width         int
	Height        int
}

// Images is a slice of Image objects
type Images []Image

// Len returns the length of the images slice
func (imgs Images) Len() int {
	return len(imgs)
}

func (imgs Images) Less(i int, j int) bool {
	return imgs[i].LastModified.Before(imgs[j].LastModified)
}

func (imgs Images) Swap(i int, j int) {
	imgs[i], imgs[j] = imgs[j], imgs[i]
}
