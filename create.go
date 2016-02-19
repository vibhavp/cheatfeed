package main

import (
	"log"
	"net/http"
	"sync"

	tf2rcon "github.com/TF2Stadium/TF2RconWrapper"
	"github.com/gorilla/websocket"
	"golang.org/x/net/xsrftoken"
)

var (
	lock      = new(sync.Mutex)
	rconAddrs = make(map[string]*tf2rcon.TF2RconConnection)
	upgrader  = websocket.Upgrader{CheckOrigin: func(_ *http.Request) bool { return true }}
)

func newSource(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	token := values.Get("xsrf-token")
	if !xsrftoken.Valid(token, key, r.RemoteAddr, "create") {
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}

	addr := values.Get("addr")
	password := values.Get("password")
	if addr == "" || password == "" {
		http.Error(w, "missing value", http.StatusBadRequest)
		return
	}

	rcon, err := tf2rcon.NewTF2RconConnection(addr, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lock.Lock()
	rconAddrs[addr] = rcon
	lock.Unlock()

	url := r.URL
	values.Del("xsrf-token")
	values.Del("password")
	url.RawQuery = values.Encode()
	url.Path = "logs"
	log.Println(url.String())

	http.Redirect(w, r, url.String(), http.StatusTemporaryRedirect)
}

func ws(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	addr := values.Get("addr")

	if addr == "" {
		log.Println(values)
		http.Error(w, "missing value", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lock.Lock()
	defer lock.Unlock()
	rcon, ok := rconAddrs[addr]
	if !ok {
		http.Error(w, "404", http.StatusNotFound)
		return
	}

	delete(rconAddrs, addr)
	h := &handler{conn}
	listener.AddSource(&tf2rcon.EventListener{
		PlayerConnected:      h.PlayerConnected,
		PlayerDisconnected:   h.PlayerDisconnected,
		PlayerGlobalMessage:  h.PlayerGlobalMessage,
		PlayerTeamMessage:    h.PlayerTeamMessage,
		PlayerSpawned:        h.PlayerSpawned,
		PlayerClassChanged:   h.PlayerClassChange,
		PlayerTeamChange:     h.PlayerTeamChange,
		PlayerKilled:         h.PlayerKilled,
		PlayerDamaged:        h.PlayerDamaged,
		PlayerHealed:         h.PlayerHealed,
		PlayerKilledMedic:    h.PlayerKilledMedic,
		PlayerUberFinished:   h.PlayerUberFinished,
		PlayerBlockedCapture: h.PlayerBlockedCapture,
		PlayerItemPickup:     h.PlayerItemPickup,
		GameOver:             h.GameOver,
		WorldRoundWin:        h.WorldRoundWin,
		TeamScoreUpdate:      h.TeamScoreUpdate,
		TeamPointCapture:     h.TeamPointCapture,
	}, rcon)

}
