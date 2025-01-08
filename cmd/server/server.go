package server

import (
	"log"
	"todo-level-5/pkg/api/handlers"
	"todo-level-5/pkg/di"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	engine   *gin.Engine
	handlers *handlers.Handlers
}

func NewHttpServer() *HttpServer {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	return &HttpServer{
		engine:   engine,
		handlers: di.InjectHandler()}
}

func (s *HttpServer) Run(port string) error {

	err := s.engine.Run(port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return err
	}
	return nil
}
