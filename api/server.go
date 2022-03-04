package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type VerifiServer interface {
	FormSubmittedWebhook(c *gin.Context)
}

type Server struct {
	logger *zap.SugaredLogger
}

func NewServer(logger *zap.SugaredLogger) VerifiServer {
	return Server{
		logger: logger,
	}
}

func (s Server) FormSubmittedWebhook(c *gin.Context) {
	s.logger.Infow("Received Form Submission")
	c.JSON(200, gin.H{"message": "form sent to server"})

	type Request struct {
		FormId       string    `json:"form_id"`
		FormName     string    `json:"form_name"`
		FormUrl      string    `json:"form_url"`
		ResponseDate time.Time `json:"response_date"`
		ResponseId   string    `json:"response_id"`
		Email        string    `json:"email"`
		FirstName    string    `json:"first_name"`
		LastName     string    `json:"last_name"`
		Pronouns     string    `json:"pronouns"`
		Major        string    `json:"major"`
		Affiliation  string    `json:"affiliation"`
		DiscordName  string    `json:"discord_name"`
		DiscordId    string    `json:"discord_id"`
	}

	var request Request
	c.Bind(&request)

	s.logger.Debugw("parsed request", "request", request)
}
