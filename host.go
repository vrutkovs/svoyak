package main

import "net/http"

func (s *ServerSettings) prepareSession(w *http.ResponseWriter) {
	// generate random UUID for session
	// generate QR code
	// encode QR code in base64 and return html page with image
}
