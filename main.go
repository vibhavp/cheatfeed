package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	webPort  = flag.String("port", "8080", "port to host web app on")
	logsPort = flag.String("logsport", "8081", "port to redirect logs to")
)

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile)
	setupRoutes()
	startListener()
	log.Printf("hosting on :%s", *webPort)
	log.Fatal(http.ListenAndServe(":"+*webPort, http.DefaultServeMux))
}
