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

	mux := my_app.NewMux()
	res := httptest.NewRecorder()  // mock
	req, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(res, req)

	assert.Equal(res.Code, 200)
	assert.Equal(res.Body.String(), "Hello World!")
}

func TestWelcomeByNameHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	mux := my_app.NewMux()
	res := httptest.NewRecorder()  // mock
	req, _ := http.NewRequest("GET", "/hello", nil)
	mux.ServeHTTP(res, req)

	assert.Equal(res.Code, 200)
	assert.Equal(res.Body.String(), "Hello World!")
}

func TestWelcomeByNameHandler_WithName(t *testing.T) {
	assert := assert.New(t)

	mux := my_app.NewMux()
	res := httptest.NewRecorder()  // mock
	req, _ := http.NewRequest("GET", "/hello?name=Mark", nil)
	mux.ServeHTTP(res, req)

	assert.Equal(res.Code, 200)
	assert.Equal(res.Body.String(), "Hello Mark!")
}
