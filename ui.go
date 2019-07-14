package main

import (
	"fmt"
	"net/http"
)

func renderStart(id string, w *http.ResponseWriter, s *ServerSettings) {
	// Return start page
	fmt.Fprintf(*w, `
	<html>
	 <head>
	  <title>Svoyak</title>
	 </head>
	 <body>
	  <a href="%s/join?id=%s">Join game</a><br/>
		QR code goes here
	 </body>
	</html>
	`, id, s.url)
}
