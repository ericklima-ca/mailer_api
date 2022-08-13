package mailer

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	mock1 := "from@email.com"
	mock2 := "From Sender <from@email.com>"
	email1, fullAddr1 := parseAddr(mock1)
	email2, fullAddr2 := parseAddr(mock2)
	log.Println(email2, fullAddr2)
	assert.Equal(t, "from@email.com", email1)
	assert.Equal(t, "from@email.com", email2)
	assert.NotEqual(t, "From Sender <from@email.com>", fullAddr1)
	assert.Equal(t, "From Sender <from@email.com>", fullAddr2)
}
