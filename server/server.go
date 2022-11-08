package server

import (
	"log"
	"net/http"
	"os"
	"syscall"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Log      *log.Logger
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

func NewServer(router *gin.Engine, log *log.Logger, shutdown chan os.Signal) *Server {
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

func (s *Server) HandleHealthCheck() gin.HandlerFunc {
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

	v1 := s.Router.Group("/v1")
	{
		v1.GET("/health-check", s.HandleHealthCheck())
	}

	s.Router.Static("/swagger", "./swagger")

}
