package http_server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	MailerService mailerservice
}

type payload struct {
	To      string `json:"to", binding:"required"`
	From    string `json:"from", binding:"required"`
	Subject string `json:"subject", binding:"required"`
	Body    string `json:"body", binding:"required"`
	Token   string `json:"token", binding:"required"`
}

type mailerservice interface {
	SendMail([]byte) error
}

func NewServer(httpServer HTTPServer) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1/api")

	v1.POST("/sendmail", httpServer.HandlerEmail)
	return router
}

func (h HTTPServer) HandlerEmail(c *gin.Context) {
	var payload payload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jsonMessage, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err := h.MailerService.SendMail(jsonMessage); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "email sent successfully",
	})
	return
}
