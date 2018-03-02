package main

import (
	"flag"
	"log"
)

var (
	dirs []string
	bind = flag.String("host", "0.0.0.0:8080", "HTTP bind address")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	flag.Parse()

	gallery, err := NewGallery(flag.Args())
	if err != nil {
		log.Fatalf("error creating gallery: %v", err)
	}

	gallery.Serve(*bind)
}
