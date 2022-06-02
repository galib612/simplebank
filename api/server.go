package api

import (
	db "github.com/galib612/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking services.
type Server struct {
	store  *db.Store
	router *gin.Engine // This router will help to send each API request to correct handler for processing
}

// Newserver creates a new http server and setup routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// VVI-- add routes to the router
	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// Start runs the Http server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
