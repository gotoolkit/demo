package main

import (
	"html/template"
	"log"
	"net/http"

	"goprojects/chat"
)

var index = template.Must(template.ParseFiles("./index.html"))

func home(w http.ResponseWriter, r *http.Request) {
	index.Execute(w, nil)
}

func main() {
	go chat.DefaultHub.Start()

	http.HandleFunc("/", home)
	http.HandleFunc("/notify", chat.Notify)
	http.HandleFunc("/ws", chat.WSHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
