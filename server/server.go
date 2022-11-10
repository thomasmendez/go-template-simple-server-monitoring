package server

import (
	"net/http"
	"os"
	"syscall"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	Log      *zap.Logger
	Router   *gin.Engine
	Shutdown chan os.Signal
}

type HealthCheck struct {
	Status string `json:"status,omitempty"`
	Host   string `json:"host,omitempty"`
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func NewServer(router *gin.Engine, log *zap.Logger, shutdown chan os.Signal) *Server {
	return &Server{
		Router:   router,
		Log:      log,
		Shutdown: shutdown,
	}
}

func (s *Server) Run(serverAddress string) {
	s.Router.Run(serverAddress)
}

func (s *Server) SignalShutdown() {
	s.Shutdown <- syscall.SIGTERM
}

func (s *Server) healthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		host, err := os.Hostname()
		if err != nil {
			host = "unavailable"
		}
		data := HealthCheck{
			Status: "up",
			Host:   host,
		}
		c.JSON(http.StatusOK, data)
	}
}

func (s *Server) SetRoutes() {

	s.Router.Use(gin.Recovery())

	route := s.Router.Group("/api")

	v1 := route.Group("/v1")
	{
<<<<<<< HEAD
		v1.GET("/health-check", s.HandleHealthCheck())
=======
		v1.GET("/health-check", s.healthCheck())
>>>>>>> template/main
	}

	route.Static("/swagger", "./swagger")

}
