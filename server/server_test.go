package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func mockZap(t *testing.T) *zap.Logger {
	return zaptest.NewLogger(t)
}

func TestHealthCheck(t *testing.T) {

	// arrange
	route := "/api/v1/health-test"
	expectedStatusCode := http.StatusOK

	srv := &Server{Log: mockZap(t)}
	r := gin.Default()
	r.GET(route, srv.healthCheck())

	// act
	req, _ := http.NewRequest("GET", route, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// assert
	assert.Equal(t, expectedStatusCode, w.Result().StatusCode)
}
