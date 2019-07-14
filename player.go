package main

import (
	"github.com/gin-gonic/gin"
)

func (s *ServerSettings) joinGame(c *gin.Context) {
	// generate random UUID for session
	// generate QR code
	// encode QR code in base64 and return html page with image
}

// use websocket to block input from other players when button is pressed?

func (s *ServerSettings) buttonClick(c *gin.Context) {
	// get current session
	// attempt to answer
}
