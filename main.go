package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
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
	url := os.Getenv("URL")
	if len(url) < 0 {
		log.Fatal("No URL env var set")
	}
	server := &ServerSettings{url: url}

	r := gin.Default()
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)

	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/health"),
		gin.Recovery(),
	)

	r.GET("/", server.prepareSession)
	r.GET("/health", health)
	r.GET("/session/:id", server.showStatus)
	r.GET("/join/:id", server.joinGame)
	r.GET("/button/:id", server.buttonClick)

	log.Printf("Server started, URL: %s\n", url)
	r.Run()
}
