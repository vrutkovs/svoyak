package main

import (
	"fmt"
	"net/http"
)

// ServerSettings stores info about the server
type ServerSettings struct {
	url string
}

func getParam(r *http.Request, key string) (string, error) {
	keys, ok := r.URL.Query()[key]
	if !ok || len(keys[0]) < 1 {
		return "", fmt.Errorf("Url Param '%s' is missing", key)
	}
	return keys[0], nil
}
