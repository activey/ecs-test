package http

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
	port int
}

func NewServerConfig() ServerConfig {
	return ServerConfig{port: 8664}
}

type RouteProvider interface {
	Route() *gin.Engine
}

type Server struct {
	route *gin.Engine
	port  int
}

func NewHttpServer(config ServerConfig) (*Server, RouteProvider) {
	route := gin.Default()
	server := &Server{
		port:  config.port,
		route: route,
	}
	return server, server
}

func (s *Server) Route() *gin.Engine {
	return s.route
}

func (s *Server) Start() {
	log.Debugf("Starting REST server on port: %d", s.port)

	err := s.route.Run(fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatal(err)
	}
}
