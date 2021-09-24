package bridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorNotification_Notify(t *testing.T) {
	emailSender := NewEmailMsgSender([]string{"test@test.com"})
	n1 := NewErrorNotification(emailSender)
	err := n1.Notify("test msg") //广播email
	assert.Nil(t, err)

	phoneSender := NewPhoneMsgSender([]string{"test@test.com"})
	n2 := NewErrorNotification(phoneSender)
	err = n2.Notify("test msg") //广播msg
	assert.Nil(t, err)
}
