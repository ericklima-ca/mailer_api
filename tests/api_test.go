package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ericklima-ca/mailer_api/http_server"
	"github.com/ericklima-ca/mailer_api/models"
	"github.com/stretchr/testify/assert"
)

func TestMailService(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/v1/api/sendmail", strings.NewReader(`
		{
			"to":      "to@email.com",
			"from":    "from@email.com",
			"subject": "test subject",
			"body":    "test body",
			"token":   "abc123"
		}
	`))
	if err != nil {
		log.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	mock := http_server.NewServer(http_server.HTTPServer{
		MailerService: models.MailerMock{},
	})

	mock.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
