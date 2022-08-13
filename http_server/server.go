package http_server

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var isEmail validator.Func = func(fl validator.FieldLevel) bool {
	email, ok := fl.Field().Interface().(string)
	if ok {
		emailPattern := regexp.MustCompile(`\s*\b[^@\s]+@[^\s]+\b\s*`)
		return emailPattern.Match([]byte(email))
	}
	return false
}

type HTTPServer struct {
	MailerService mailerservice
}

type payload struct {
	To      string `json:"to" binding:"required,isEmail"`
	From    string `json:"from" binding:"required,isEmail"`
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
	Token   string `json:"token" binding:"required"`
}

type mailerservice interface {
	SendMail([]byte) error
}

func NewServer(httpServer HTTPServer) *gin.Engine {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("isEmail", isEmail)
	}
	v1 := router.Group("/v1/api")

	v1.POST("/sendmail", httpServer.HandlerEmail)
	v1.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "v1")
	})
	return router
}

func (h HTTPServer) HandlerEmail(c *gin.Context) {
	var payload payload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "invalid payload",
		})
		return
	}

	jsonMessage, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err := h.MailerService.SendMail(jsonMessage); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"msg":   "email not sent",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "email sent successfully",
	})
}

// type Booking struct {
// 	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
// 	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
// }
