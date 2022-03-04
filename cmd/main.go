package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"verifi/api"
)

func main() {
	c := config{
		url:  "",
		port: "8080",
	}

	port := os.Getenv("PORT")
	if port != "" {
		c.port = port
	}

	url := os.Getenv("URL")
	if url != "" {
		c.url = url
	}

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()

	s := api.NewServer(sugar)

	r := gin.New()

	api := r.Group("/api")
	api.POST("/form", s.FormSubmittedWebhook)

	r.Run(fmt.Sprintf("%s:%s", c.url, c.port))
}

type config struct {
	url  string
	port string
}
