package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *ServerSettings) joinGame(c *gin.Context) {
	// TODO: generate QR code
	c.HTML(http.StatusOK, "templates/join-session.tmpl", gin.H{
		"ID": c.Param("id"),
	})
}

func (s *ServerSettings) buttonClick(c *gin.Context) {
	// get current session
	// attempt to answer
}
