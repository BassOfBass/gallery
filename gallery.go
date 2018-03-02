package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	tmpl = template.Must(template.New("index").Parse(templateStr))
)

// Gallery is a set of Albums served over HTTP
type Gallery struct {
	Albums map[string]Album
}

// NewGallery returns a new Gallery
func NewGallery(dirs []string) (*Gallery, error) {
	if dirs == nil || len(dirs) == 0 {
		return nil, fmt.Errorf("no directories specified to serve")
	}

	g := &Gallery{Albums: map[string]Album{}}
	for _, d := range dirs {
		a, err := NewAlbum(d)
		if err != nil {
			return nil, fmt.Errorf("error creating album for %s: %v", d, err)
		}

		g.Albums[a.Name] = a
	}

	return g, nil
}

// IndexHandler handles requests to the gallery homepage
func (g *Gallery) IndexHandler(w http.ResponseWriter, r *http.Request) {
	albums := []Album{}
	for _, a := range g.Albums {
		albums = append(albums, a)
	}
	buf := &bytes.Buffer{}
	err := tmpl.Execute(buf, struct{ Albums []Album }{albums})
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(buf.Bytes())
}

// ImageHandler handles requests for images
func (g *Gallery) ImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	a, ok := g.Albums[vars["album"]]
	if !ok {
		http.Error(w,
			fmt.Sprintf("%s: unknown album '%s'", http.StatusText(http.StatusBadRequest), vars["album"]),
			http.StatusBadRequest)
		return
	}
	for _, img := range a.Images {
		if vars["id"] == img.Md5 {
			http.ServeFile(w, r, img.Path)
			return
		}
		if vars["id"] == img.ThumbnailMd5 {
			http.ServeFile(w, r, img.ThumbnailPath)
			return
		}
	}
	http.Error(w, fmt.Sprintf("%s: unknown image '%s'", http.StatusText(http.StatusBadRequest), vars["id"]),
		http.StatusBadRequest)
}

// Serve starts a HTTP server for the gallery
func (g *Gallery) Serve(addr string) {
	r := mux.NewRouter()

	r.HandleFunc("/", g.IndexHandler)
	r.HandleFunc("/{album}/{id}", g.ImageHandler)

	log.Printf("Listening on %v...", addr)
	log.Fatal(http.ListenAndServe(addr, handlers.CombinedLoggingHandler(os.Stdout, r)))
}
