package api

import (
	"io"
	"net/http"
	"os"
	"strings"
)

var here = os.Getenv("GOPATH") + "/src/goprojects/api/images/"

// UploadImage allows us to upload an image.
func UploadImage(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
		errJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	out, err := os.Create(here + header.Filename)
	if err != nil {
		errJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		errJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]string{
		"filename": header.Filename,
	})
}

// ShowImage shows the image based on the filename found in the path
func ShowImage(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimLeft(r.URL.Path, "/image/")
	file, err := os.Open(here + name)
	if err != nil {
		errJSON(w, err.Error(), http.StatusNotFound)
		return
	}

	buf := pool.Get()
	defer pool.Put(buf)
	_, err = io.Copy(buf, file)
	if err != nil {
		errJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	buf.WriteTo(w)
}
