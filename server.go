package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

var indexTemplate = template.Must(template.ParseFiles("template/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	alerts := GetStatus()
	err := indexTemplate.Execute(&buf, alerts)
	if err != nil {
		log.Println("Index template rendering failed: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(buf.Bytes())
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	alerts := GetStatus()
	json.NewEncoder(w).Encode(alerts)
}
