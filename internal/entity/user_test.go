package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestNewUser(t *testing.T) {
	user, err := NewUser("victory", "j@j.com", "123456")
	assert.Nil(t, err)
	assert.Equal(t, "victory", user.Name)
	assert.Equal(t, "j@j.com", user.Email)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.NotNil(t, user)

}

func TestUser_ValidatePassswordUser(t *testing.T) {
	user, err := NewUser("victory", "j@j.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePasswordUser("123456"))
	assert.False(t, user.ValidatePasswordUser("123457"))
	assert.NotEqual(t, "123456", user.Password)
}