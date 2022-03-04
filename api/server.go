package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
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


	body, _ := ioutil.ReadAll(c.Request.Body)
	s.logger.Debugw("printing body", "body", body)
}
