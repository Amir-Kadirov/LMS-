package pkg

import (
	"backend_course/lms/pkg/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendingEmail(t *testing.T) {
	err:=smtp.SendMail("devamirkadirov@gmail.com","Hello world")
	assert.NoError(t,err)
}