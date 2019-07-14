package main

import (
	"net/http"

	"github.com/rs/xid"
)

func (s *ServerSettings) prepareSession(w *http.ResponseWriter) {
	// generate random UUID for session
	// generate QR code
	// encode QR code in base64 and return html page with image
	renderStart(xid.New().String(), w, s)
}
