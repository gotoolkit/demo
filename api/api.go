package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gotoolkit/demo/api/cms"
)

var pool = New()

// Doc lists all the routes for our API
func Doc(w http.ResponseWriter, r *http.Request) {
	data := (map[string]string{
		"all_pages_url":   "/pages",
		"page_url":        "/pages/{id}",
		"create_page_url": "/newpage",
	})
	writeJSON(w, data)
}

// AllPages returns all the pages
func AllPages(w http.ResponseWriter, r *http.Request) {
	data, err := cms.GetPages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, data)
}

// GetPage gets a single page from the API
func GetPage(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimLeft(r.URL.Path, "/pages/")
	data, err := cms.GetPage(id)
	if err != nil {
		errJSON(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, data)
}

// CreatePage creates a new post or page
func CreatePage(w http.ResponseWriter, r *http.Request) {
	page := new(cms.Page)
	err := json.NewDecoder(r.Body).Decode(page)
	if err != nil {
		errJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := cms.CreatePage(page)
	if err != nil {
		errJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]int{
		"user_id": id,
	})
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	buf := pool.Get()
	defer pool.Put(buf)
	err := json.NewEncoder(buf).Encode(data)
	if err != nil {
		errJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	buf.WriteTo(w)
}

func errJSON(w http.ResponseWriter, err string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte("{\n\terror: " + err + "\n}\n"))
}
