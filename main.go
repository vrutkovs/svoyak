package main

import (
	"log"
	"net/http"
	"os"
)

//TODO: use gorilla mux insted of arg parsing

// handleStart would display a random QR code for the session
func (s *ServerSettings) handleStart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.prepareSession(&w)
	}
}

// handleJoin would add user to the session
func (s *ServerSettings) handleJoin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Verify session is present
		session, err := getParam(r, "id")
		if err != nil {
			return
		}
		s.joinGame(&w, session)
	}
}

// handleStart would display a random QR code for the session
func (s *ServerSettings) handleButton() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Verify session is present
		session, err := getParam(r, "id")
		if err != nil {
			return
		}
		s.click(&w, session)
	}
}

func main() {
	url := os.Getenv("URL")
	if len(url) < 0 {
		log.Fatal("No URL env var set")
	}
	server := &ServerSettings{url: url}
	log.Printf("Server started, URL: %s\n", url)

	http.HandleFunc("/", server.handleStart())
	http.HandleFunc("/join", server.handleJoin())
	http.HandleFunc("/button", server.handleButton())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
