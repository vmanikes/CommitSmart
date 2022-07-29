package http

import (
	v1 "BankingService/internal/handler/http/v1"
	"github.com/flannel-dev-lab/cyclops/v2/middleware"
	"github.com/flannel-dev-lab/cyclops/v2/response"
	"github.com/flannel-dev-lab/cyclops/v2/router"
	"net/http"
)

func GetRoutes(handler *v1.Handler) *router.Router {
	routerObj := router.New(false, nil, nil)

	routerObj.Get("/health", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response.SuccessResponse(http.StatusOK, writer, nil)
		return
	}))

	routerObj.Post("/api/v1/deposit", middleware.NewChain().Then(handler.Deposit))
	routerObj.Post("/api/v1/withdraw", middleware.NewChain().Then(handler.Withdraw))
	routerObj.Post("/api/v1/transfer", middleware.NewChain().Then(handler.Transfer))
	routerObj.Get("/api/v1/users/:user_id/transactions", middleware.NewChain().Then(handler.Transactions))

	return routerObj
}
