package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

// prepareSession would display a random QR code for the session
func (s *ServerSettings) prepareSession(c *gin.Context) {
	// generate random UUID for session
	sessionPage := fmt.Sprintf("%s/session/%s", s.url, xid.New().String())
	// redirect to session page
	c.Redirect(http.StatusMovedPermanently, sessionPage)
}

func (s *ServerSettings) showStatus(c *gin.Context) {
	// TODO: generate QR code
	c.HTML(http.StatusOK, "/templates/session.tmpl", gin.H{
		"Url": s.url,
		"ID":  c.Param("id"),
	})
}
