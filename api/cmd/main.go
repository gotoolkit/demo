package main

import (
	"log"
	"net/http"
	"os"

	"goprojects/api"
)

func main() {
	os.Mkdir("images", 0777)

	http.HandleFunc("/", api.Doc)
	http.HandleFunc("/image/", api.ShowImage)
	http.HandleFunc("/newpage", api.CreatePage)
	http.HandleFunc("/pages", api.AllPages)
	http.HandleFunc("/pages/", api.GetPage)
	http.HandleFunc("/upload", api.UploadImage)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
