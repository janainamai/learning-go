package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Janaina Mai", "jana@hotmail.com", "12345")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Janaina Mai", user.Name)
	assert.Equal(t, "jana@hotmail.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Janaina Mai", "jana@hotmail.com", "12345")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("12345"))
	assert.False(t, user.ValidatePassword("123456"))
	assert.NotEqual(t, "12345", user.Password)
}
