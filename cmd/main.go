package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"verifi/api"
)

func main() {
	c := config{
		url:  "localhost",
		port: "8080",
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
