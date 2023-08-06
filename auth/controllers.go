package auth

import (
	"auth-and-gateway-microservice/custom_errors"
	db "auth-and-gateway-microservice/db/sqlc"
	"auth-and-gateway-microservice/utils"
	"auth-and-gateway-microservice/validate"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	store *db.Store
}

func NewController(store *db.Store) *Controller {
	return &Controller{store: store}
}

func (ac *Controller) SignUpUser(ctx *gin.Context) {
	var payload *db.CreateUserTxParams
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(utils.INVALID_DATA, err.Error()))
		return
	}
	err := validate.CreateUserTxParams(payload)
	if err != nil {
		tErr := custom_errors.TranslateError(err)
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(tErr.Code, tErr.Error()))
		return
	}
	err = ac.store.TxCreateUser(ctx, payload, "ORDINARY")
	if err != nil {
		tErr := custom_errors.TranslateError(err)
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(tErr.Code, tErr.Error()))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "success"})
}

func (ac *Controller) SignInUser(ctx *gin.Context) {
	config, _ := utils.LoadConfig(".")
	var payload *SignInInputSerializer
	var loginType string
	var login string
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(utils.INVALID_DATA, err.Error()))
		return
	}
	switch {
	case payload.Email != "" && payload.Username == "":
		loginType = "EMAIL"
		login = payload.Email
	case payload.Username != "" && payload.Email == "":
		loginType = "USERNAME"
		login = payload.Username
	case payload.Username != "" && payload.Email != "":
		loginType = "EMAIL"
		login = payload.Email
	}
	user, err := ac.store.GetUser(ctx, login, loginType, payload.Password)
	if err != nil {
		tErr := custom_errors.TranslateError(err)
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(tErr.Code, tErr.Error()))
		return
	}
	userPayload := db.UserPayloadToCookie(user)
	accessToken, err := utils.CreateToken(config.JwtAccessExpiredIn, userPayload, "access", config.JwtAccessTokenPrivateKey)
	if err != nil {
		tErr := custom_errors.TranslateError(err)
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(tErr.Code, tErr.Error()))
		return
	}

	refreshToken, err := utils.CreateToken(config.JwtRefreshExpiredIn, userPayload, "refresh", config.JwtAccessTokenPrivateKey)
	if err != nil {
		tErr := custom_errors.TranslateError(err)
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(tErr.Code, tErr.Error()))
		return
	}

	ctx.SetCookie("access_token", accessToken, config.JwtAccessMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refreshToken, config.JwtRefreshMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.JwtAccessMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"msg": "success", "code": 200})
}

func (ac *Controller) RefreshAccessToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(utils.GET_REFRESH_TOKEN_FROM_COOKIES_ERR, err.Error()))
		return
	}
	config, err := utils.LoadConfig(".")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(utils.GET_CONFIG_ERR, err.Error()))
		return
	}
	sub, err := utils.ValidateToken(refreshToken, "refresh", config.JwtAccessTokenPublicKey)
	if err != nil {
		tErr := custom_errors.TranslateError(err)
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(tErr.Code, tErr.Error()))
		return
	}
	tokenPayload := sub.(map[string]interface{})
	user, err := ac.store.GetUserById(ctx, int32(tokenPayload["user_id"].(float64)))
	if err != nil {
		tErr := custom_errors.TranslateError(err)
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(tErr.Code, tErr.Error()))
		return
	}
	userPayload := db.UserPayloadToCookie(user)
	accessToken, err := utils.CreateToken(config.JwtAccessExpiredIn, userPayload, "access", config.JwtAccessTokenPrivateKey)
	if err != nil {
		tErr := custom_errors.TranslateError(err)
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(tErr.Code, tErr.Error()))
		return
	}

	ctx.SetCookie("access_token", accessToken, config.JwtAccessMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.JwtAccessMaxAge*60, "/", "localhost", false, false)
	ctx.JSON(http.StatusOK, gin.H{"msg": "success", "code": 200})
}
