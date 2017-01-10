package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/status", statusHandler)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":9001", nil))
}
