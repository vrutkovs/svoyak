package main

import (
	"fmt"
	"net/http"

	"github.com/rs/xid"
)

func (s *ServerSettings) prepareSession(w *http.ResponseWriter, r *http.Request) {
	// generate random UUID for session
	// redirect to start page
	http.Redirect(*w, r, fmt.Sprintf("%s/session?id=%s", s.url, xid.New().String()), 301)
}

func (s *ServerSettings) statusPage(w *http.ResponseWriter, session string) {
	// generate QR code
	// encode QR code in base64 and return html page with image
	renderStart(session, w, s)
}
