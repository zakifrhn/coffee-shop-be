package pkg

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

var data = claims{}
var errTest error
var id = "11111"
var role = "user"
var tokenTest = ""

func TestNewToken(t *testing.T) {
	newTest := &claims{Id: id, Role: role, RegisteredClaims: jwt.RegisteredClaims{Issuer: "backendgolang"}}
	testToken := NewToken(id, role)
	testToken.RegisteredClaims.ExpiresAt = nil
	data = *testToken
	assert.Equal(t, newTest, testToken, "data ga ada")
}

func TestGenerate(t *testing.T) {

	test, errTest := data.Generate()
	tokenTest = test

	assert.NotEqual(t, "", test, "token kosong")
	assert.Nil(t, errTest, "Generate token error")
}

func TestVerifyToken(t *testing.T) {
	t.Run("Verify Berhasil", func(t *testing.T) {
		expectResult := &data
		testVerify, errTest := VerifyToken(tokenTest)

		assert.Equal(t, expectResult, testVerify, "token harus sama")
		assert.Nil(t, errTest, "verify failed")
	})
}
