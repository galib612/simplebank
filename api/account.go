package api

import (
	"net/http"

	db "github.com/galib612/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR RUP"`
}

//Info -- when using gin, everything we do inside a handler will involve the context object
//        it provide lots of covenient methods to read inputs parameters and write responses.
func (server *Server) createAccount(ctx *gin.Context) {
	var req CreateAccountRequest
	if err := ctx.BindJSON(&req); err != nil {
		// Second Object is json object that want to send to client
		// So here we need a function to convert this err into key value object so that gin an
		// serialize it to JSOn before returning to the client.
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  req.Balance,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)

}
