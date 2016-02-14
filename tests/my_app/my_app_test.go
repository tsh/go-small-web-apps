package my_app_test

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"../my_app"

	"github.com/stretchr/testify/assert"
)

func TestHomePageHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()  // mock
	req, _ := http.NewRequest("GET", "/", nil)
	my_app.HomePageHandler(res, req)

	assert.Equal(res.Code, 200)
}
