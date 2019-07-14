package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rs/xid"
)

// WSMessage represents websocket message format
type WSMessage struct {
	action string
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (s *ServerSettings) handleStatusViaWS(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Printf("Failed to upgrade ws: %+v", err)
		return
	}

	s.statusWebSocket = conn
}

// sendStatusViaWebSocket would display client ID
func (s *ServerSettings) sendStatusViaWebSocket(id string) {
	// TODO: create a structure
	responseJSON := fmt.Sprintf("{'action': 'joined', 'id': '%s'}", id)
	if s.statusWebSocket != nil {
		s.statusWebSocket.WriteMessage(websocket.TextMessage, []byte(responseJSON))
	}
}

func (s *ServerSettings) handleJoinViaWS(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Printf("Failed to upgrade ws: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %s", err)
			break
		}
		if t != websocket.TextMessage {
			log.Printf("Not a text message: %d", t)
			continue
		}
		var m WSMessage
		err = json.Unmarshal(msg, &m)
		if err != nil {
			log.Printf("Failed to unmarshal message '%+v': %+v", string(msg), err)
		}
		if m.action == "join" {
			clientID := xid.New().String()
			s.handleJoin(m, t, conn, clientID)
			s.sendStatusViaWebSocket(clientID)
		}
	}
}

// handleJoin would generate a userID for a button and generate unique URL
func (s *ServerSettings) handleJoin(w WSMessage, t int, conn *websocket.Conn, id string) {
	//TODO: associate button ID with session
	buttonURL := fmt.Sprintf("/button/%s", id)
	conn.WriteMessage(t, []byte(buttonURL))
}

// prepareSession would display a random QR code for the session
func (s *ServerSettings) generateSession(c *gin.Context) {
	// generate random UUID for session
	sessionPage := fmt.Sprintf("session/%s", xid.New().String())
	log.Printf("Generated URL: %s\n", sessionPage)
	// redirect to session page
	c.Redirect(http.StatusMovedPermanently, sessionPage)
}

func (s *ServerSettings) showStatus(c *gin.Context) {
	// TODO: generate QR code
	c.HTML(http.StatusOK, "templates/session.tmpl", gin.H{
		"ID": c.Param("id"),
	})
}
