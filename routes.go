package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"text/template"

	"golang.org/x/net/xsrftoken"
)

var createForm = template.Must(template.ParseFiles("create.html"))

var key string

func init() {
	buf := make([]byte, 32)
	rand.Read(buf)
	hash := sha256.Sum256(buf)
	key = hex.EncodeToString(hash[:])
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		createForm.Execute(w, map[string]string{
			"XSRFToken": xsrftoken.Generate(key, r.RemoteAddr, "create"),
		})
	})
	http.HandleFunc("/logs.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "logs.js")
	})
	http.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "logs.html")
	})
	http.HandleFunc("/create", newSource)
	http.HandleFunc("/ws", ws)
}
