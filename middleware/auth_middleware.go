package middleware

import (
	"auth-and-gateway-microservice/custom_errors"
	db "auth-and-gateway-microservice/db/sqlc"
	"auth-and-gateway-microservice/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AuthMiddleware(store *db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessToken string
		var cookieAccessToken string
		config, err := utils.LoadConfig(".")
		if err != nil {
			log.Fatalf("could not load config: %v", err)
		}
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		accessToken, err = utils.GetAccessToken(authorizationHeader)
		if err != nil {
			cookieAccessToken, err = ctx.Cookie("access_token")
			if err != nil {
				tErr := custom_errors.TranslateError(err)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(tErr.Code, tErr.Error()))
				return
			}
			accessToken = cookieAccessToken
		}
		if err != nil {
			tErr := custom_errors.TranslateError(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(tErr.Code, tErr.Error()))
		}
		sub, err := utils.ValidateToken(accessToken, "access", config.JwtAccessTokenPrivateKey)
		if err != nil {
			tErr := custom_errors.TranslateError(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(tErr.Code, tErr.Error()))
			return
		}
		tokenPayload := sub.(map[string]interface{})
		user, err := store.GetUserById(ctx, int32(tokenPayload["user_id"].(float64)))
		if err != nil {
			tErr := custom_errors.TranslateError(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(tErr.Code, tErr.Error()))
			return
		}
		ctx.Set("User", db.UserPayloadToCookie(user))
		ctx.Next()
	}
}
