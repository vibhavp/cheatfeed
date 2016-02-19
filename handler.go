package main

import (
	tf2rcon "github.com/TF2Stadium/TF2RconWrapper"
	"github.com/gorilla/websocket"
)

type handler struct {
	*websocket.Conn
}

type event struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

func (h *handler) PlayerConnected(data tf2rcon.PlayerData) {
	h.WriteJSON(event{"PlayerConnected", data})
}

func (h *handler) PlayerDisconnected(data tf2rcon.PlayerData) {
	h.WriteJSON(event{"PlayerDisconnected", data})
}

func (h *handler) PlayerGlobalMessage(data tf2rcon.PlayerData, msg string) {
	h.WriteJSON(event{"PlayerGlobalMessage", data})
}

func (h *handler) PlayerTeamMessage(data tf2rcon.PlayerData, msg string) {
	h.WriteJSON(event{"PlayerTeamMessage", data})
}

func (h *handler) PlayerSpawned(data tf2rcon.PlayerData, class string) {
	h.WriteJSON(event{"PlayerSpawned", data})
}

func (h *handler) PlayerClassChange(data tf2rcon.PlayerData, class string) {
	h.WriteJSON(event{"PlayerClassChange", data})
}

func (h *handler) PlayerTeamChange(data tf2rcon.PlayerData, team string) {
	h.WriteJSON(event{"PlayerTeamChange", data})
}

func (h *handler) PlayerKilled(kill tf2rcon.PlayerKill) {
	h.WriteJSON(event{"PlayerKilled", kill})
}

func (h *handler) PlayerDamaged(dmg tf2rcon.PlayerDamage) {
	h.WriteJSON(event{"PlayerDamaged", dmg})
}

func (h *handler) PlayerHealed(heal tf2rcon.PlayerHeal) {
	h.WriteJSON(event{"PlayerHealed", heal})
}

func (h *handler) PlayerKilledMedic(trigger tf2rcon.PlayerTrigger) {
	h.WriteJSON(event{"PlayerKilledMedic", trigger})
}

func (h *handler) PlayerUberFinished(data tf2rcon.PlayerData) {
	h.WriteJSON(event{"PlayerUberFinished", data})
}

func (h *handler) PlayerBlockedCapture(cp tf2rcon.CPData, data tf2rcon.PlayerData) {
	h.WriteJSON(event{"PlayerBlockedCapture", struct {
		tf2rcon.CPData
		tf2rcon.PlayerData
	}{cp, data}})
}

func (h *handler) PlayerItemPickup(item tf2rcon.ItemPickup) {
	h.WriteJSON(event{"PlayerItemPickup", item})
}

func (h *handler) GameOver() {
	h.WriteJSON(event{"GameOver", ""})
}

func (h *handler) WorldRoundWin(team string) {
	h.WriteJSON(event{"WorldRoundWin", struct {
		Team string `json:"team"`
	}{team}})
}

func (h *handler) TeamScoreUpdate(team tf2rcon.TeamData) {
	h.WriteJSON(event{"TeamScoreUpdate", team})
}

func (h *handler) TeamPointCapture(team tf2rcon.TeamData) {
	h.WriteJSON(event{"WorldRoundWin", team})
}
