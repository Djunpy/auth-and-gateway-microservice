package accounts

import (
	db "auth-and-gateway-microservice/db/sqlc"
	"auth-and-gateway-microservice/middleware"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	controller *Controller
	store      *db.Store
}

func NewRoutes(controller *Controller, store *db.Store) Routes {
	return Routes{controller: controller, store: store}
}

func (ar *Routes) AccountsRoute(rg *gin.RouterGroup) {
	router := rg.Group("accounts")
	router.POST("/add-address", middleware.AuthMiddleware(ar.store), ar.controller.AddUserAddress)
	router.PUT("/update-address", middleware.AuthMiddleware(ar.store), ar.controller.UpdateUserAddress)
	router.POST("/add-phone", middleware.AuthMiddleware(ar.store), ar.controller.AddUserPhone)
	router.PUT("/update-phone", middleware.AuthMiddleware(ar.store), ar.controller.UpdateUserPhone)
	router.POST("/update", middleware.AuthMiddleware(ar.store), ar.controller.UpdateAccount)
}
