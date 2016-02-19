package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"text/template"

	"golang.org/x/net/xsrftoken"
)

var createForm = template.Must(template.New("form").Parse(`
<html>
  <body>
    Automatic Setup
    <form action="/create" method="get">
      RCON Address:<br>
      <input type="text" name="addr"><br>
      Password:<br>
      <input type="text" name="password"><br><br>
      <input type="hidden" name="xsrf-token" value="{{.XSRFToken}}">
      <input type="submit" value="Submit">
    </form>
    Manual Setup (WIP)
  </body>
</html>
`))

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
