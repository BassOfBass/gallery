package main

var (
	// AllowedExtensions are allowed to be served
	AllowedExtensions = []string{".gif", ".jpeg", ".jpg", ".png"}

	// ThumbnailDimensions are the maximum dimensions of a thumbnail
	ThumbnailDimensions = struct {
		Width  uint
		Height uint
	}{
		Width:  256,
		Height: 256,
	}
)
