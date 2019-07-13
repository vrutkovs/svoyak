package main

import "net/http"

func (s *ServerSettings) joinGame(w *http.ResponseWriter, session string) {
	// generate random UUID for session
	// generate QR code
	// encode QR code in base64 and return html page with image
}

// use websocket to block input from other players when button is pressed?

func (s *ServerSettings) click(w *http.ResponseWriter, session string) {
	// get current session
	// attempt to answer
}
