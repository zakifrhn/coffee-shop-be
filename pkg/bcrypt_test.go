package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var password = "abcd12345"
var hasedPassword string
var errors error

func TestHashPassword(t *testing.T) {
	hasedPassword, errors = HashPasword(password)
	assert.NoError(t, errors, "error ketika sedang menghasing")
	assert.NotEqual(t, password, hasedPassword, "password tidak terhasing")
}

func TestVerifyPassword(t *testing.T) {
	t.Run("verify success", func(t *testing.T) {
		var hashPassword = VerifyPassword(hasedPassword, password)
		assert.Nil(t, hashPassword, "password salah")
	})

	t.Run("verify failed", func(t *testing.T) {
		var hashPassword = VerifyPassword(hasedPassword, "kljisa")
		assert.NotNil(t, hashPassword, "password masih benar")
	})
}
