package main

import (
	"io/ioutil"
	"log"
	"net/http"

	tf2rcon "github.com/TF2Stadium/TF2RconWrapper"
)

func getlocalip() string {
	resp, err := http.Get("http://api.ipify.org")
	if err != nil {
		log.Fatal(err)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}

var (
	listener *tf2rcon.Listener
)

func startListener() {
	var err error
	listener, err = tf2rcon.NewListenerAddr(*logsPort, getlocalip()+":"+*logsPort)
	if err != nil {
		log.Fatal(err)
	}
}
