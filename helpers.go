package main

import "github.com/gorilla/websocket"

// ServerSettings stores info about the server
type ServerSettings struct {
	url             string
	statusWebSocket *websocket.Conn
}
