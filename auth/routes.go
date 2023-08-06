package auth

import (
	db "auth-and-gateway-microservice/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	controller *Controller
	store      *db.Store
}

func NewRoutes(c *Controller, store *db.Store) Routes {
	return Routes{controller: c, store: store}
}

func (r *Routes) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")
	router.POST("/sign-up", r.controller.SignUpUser)
	router.POST("/sign-in", r.controller.SignInUser)
	router.POST("/refresh-access-token", r.controller.RefreshAccessToken)
}
