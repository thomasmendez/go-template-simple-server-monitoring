package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	mockLog = log.New(ioutil.Discard, "", 0)
)

func TestHealthCheck(t *testing.T) {

	// arrange
	route := "/api/v1/health-test"
	expectedStatusCode := http.StatusOK

	srv := &Server{Log: mockLog}
	r := gin.Default()
	r.GET(route, srv.HandleHealthCheck())

	// act
	req, _ := http.NewRequest("GET", route, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// assert
	assert.Equal(t, expectedStatusCode, w.Result().StatusCode)
}
