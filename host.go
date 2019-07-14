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

func (s *ServerSettings) statusWebSocket(c *gin.Context) {
	s.wsSessionJoin(c.Writer, c.Request)
}

func (s *ServerSettings) wsSessionJoin(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("Failed to upgrade ws: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		var m WSMessage
		err = json.Unmarshal(msg, &m)
		if err != nil {
			log.Printf("Failed to unmarshal message '%+v': %+v", msg, err)
		}
		if m.action == "join" {
			s.handleJoin(m, t, conn)
		}
		// conn.WriteMessage(t, msg)
	}
}

// handleJoin would generate a userID for a button and generate unique URL
func (s *ServerSettings) handleJoin(w WSMessage, t int, conn *websocket.Conn) {
	//TODO: associate button ID with session
	buttonURL := fmt.Sprintf("%s/button/%s", s.url, xid.New().String())
	conn.WriteMessage(t, []byte(buttonURL))
}

// prepareSession would display a random QR code for the session
func (s *ServerSettings) generateSession(c *gin.Context) {
	// generate random UUID for session
	sessionPage := fmt.Sprintf("%s/session/%s", s.url, xid.New().String())
	log.Printf("Generated URL: %s\n", sessionPage)
	// redirect to session page
	c.Redirect(http.StatusMovedPermanently, sessionPage)
}

func (s *ServerSettings) showStatus(c *gin.Context) {
	// TODO: generate QR code
	c.HTML(http.StatusOK, "templates/session.tmpl", gin.H{
		"Url": s.url,
		"ID":  c.Param("id"),
	})
}
