package http_server

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ericklima-ca/mailer_api/models"
	"github.com/stretchr/testify/assert"
)

func TestVersionRoute(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/v1/api/version", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	mock := NewServer(HTTPServer{
		MailerService: models.MailerMock{},
	})

	mock.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "v1", w.Body.String())
}
