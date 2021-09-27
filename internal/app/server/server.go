package server

import (
	"log"

	routes "github.com/vitorbgouveia/go-rabbitmq/internal/app/router"

	"github.com/gin-gonic/gin"
)

// Server configs
type Server struct {
	port   string
	server *gin.Engine
}

// NewServer create server
func newServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

// Init server
func Init() {
	s := newServer()

	go func() {
		router := routes.ConfigRoutes(s.server)

		log.Print("Server is running at port: ", s.port)
		log.Fatal(router.Run(":" + s.port))
	}()
}
