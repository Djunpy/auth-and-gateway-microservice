package utils

import "github.com/gin-gonic/gin"

const (
	// config custom_errors 35-40
	GET_CONFIG_ERR = 35

	// cookie 20-35
	GET_REFRESH_TOKEN_FROM_COOKIES_ERR = 20

	// post request custom_errors 5-20
	INVALID_DATA = 5

	// users custom_errors 60-90
	INCORRECT_PASSWORD = 60
)

//const (
//	SUCCESS               = 200
//	ERROR                 = 500
//	INVALID_PARAMS        = 400
//	INVALID_DATA          = 1
//	INVALID_LOGIN_TYPE    = 30
//	INCORRECT_PASSWORD    = 31
//	TOKEN_ERROR           = 50
//	GET_REFRESH_TOKEN_ERR = 51
//	GET_CONFIG_ERR        = 70
//)

func ErrorResponse(code int, msg string) gin.H {
	return gin.H{"code": code, "msg": msg}
}
