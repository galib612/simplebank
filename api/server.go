package api

import (
<<<<<<< HEAD
	authdb "github.com/galib612/simplebank/authdb/sqlc"
=======
>>>>>>> dace6cc8a3810b11427e67eccada66a80cd49e66
	db "github.com/galib612/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking services.
type Server struct {
<<<<<<< HEAD
	postgresStore db.Store
	authdbStore   authdb.Store
	router        *gin.Engine // This router will help to send each API request to correct handler for processing
}

// Newserver creates a new http server and setup routing.
func NewServer(postgresStore db.Store, authdbStore authdb.Store) *Server {
	server := &Server{postgresStore: postgresStore,
		authdbStore: authdbStore}

	router := gin.Default()

	// VVI-- add routes to the router
	router.POST("/accounts/create/", server.createAccount)
	router.GET("/accounts/get/:id", server.getAccount)
	router.GET("/accounts/get/", server.listAccount)
	router.POST("accounts/update/", server.updateAccount)
	router.GET("accounts/delete/:id", server.deleteAccount)

	// Login Authentication Api
	router.POST("/user/signup/", server.createUser)
	router.POST("/user/login/", server.loginUser)
=======
	store  *db.Store
	router *gin.Engine // This router will help to send each API request to correct handler for processing
}

// Newserver creates a new http server and setup routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// VVI-- add routes to the router
	router.POST("/accounts", server.createAccount)
>>>>>>> dace6cc8a3810b11427e67eccada66a80cd49e66

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
