package pkg

import (
	"inter/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var codes int = 500
var err error
var status string

type ResponseMock struct {
	mock.Mock
}

func TestGetStatus(t *testing.T) {
	expectResult := ""
	status = getStatus(codes)
	assert.NotEqual(t, expectResult, getStatus(codes), "No description")
}

func TestNewRes(t *testing.T) {
	expectResult := &Response{Code: codes, Status: status}

	testRes := NewRes(codes, &config.Result{})

	assert.Equal(t, expectResult, testRes, "Gak sama")
}
