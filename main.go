package main

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for _, name := range AssetNames() {
		file, err := AssetInfo(name)
		if err != nil || file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		contents, err := Asset(name)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(contents))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

// health is k8s endpoint for liveness check
func health(c *gin.Context) {
	c.String(http.StatusOK, "")
}

func main() {
	server := &ServerSettings{}

	r := gin.New()
	// Don't log k8s health endpoint
	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/health"),
		gin.Recovery(),
	)
	r.GET("/health", health)

	// Load templates from bin assets
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)

	r.GET("/", server.generateSession)
	r.GET("/session/:id", server.showStatus)
	r.GET("/join/:id", server.joinGame)
	r.GET("/button/:id", server.buttonClick)

	r.GET("/ws/session-status/:id", server.handleStatusViaWS)
	r.GET("/ws/join-session/:id", server.handleJoinViaWS)
	r.Run()
}
