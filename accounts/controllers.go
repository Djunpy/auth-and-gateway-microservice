package accounts

import (
	"auth-and-gateway-microservice/custom_errors"
	db "auth-and-gateway-microservice/db/sqlc"
	"auth-and-gateway-microservice/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	store *db.Store
}

func NewController(store *db.Store) *Controller {
	return &Controller{store: store}
}

func (ac *Controller) AddUserAddress(ctx *gin.Context) {
	var payload *UserAddress
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(utils.INVALID_DATA, err.Error()))
		return
	}
	currentUser, _ := ctx.MustGet("User").(db.UserPayload)
	addressArgs := &db.CreateUserAddressParams{
		UserID:     currentUser.UserId,
		City:       payload.City,
		Street:     sql.NullString{String: payload.Street, Valid: true},
		PostalCode: sql.NullInt64{Int64: payload.PostalCode, Valid: true},
	}
	address, err := ac.store.CreateUserAddress(ctx, *addressArgs)
	if err != nil {
		tErr := custom_errors.TranslateError(err)
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(tErr.Code, tErr.Error()))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "payload": ResponseUserAddress(address)})
}

func (ac *Controller) UpdateUserAddress(ctx *gin.Context) {
	var payload *UserAddress
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(utils.INVALID_DATA, err.Error()))
		return
	}
	currentUser, _ := ctx.MustGet("User").(db.UserPayload)
	addressArgs := &db.UpdateUserAddressParams{
		UserID:     currentUser.UserId,
		City:       sql.NullString{String: payload.City, Valid: true},
		Street:     sql.NullString{String: payload.Street, Valid: true},
		PostalCode: sql.NullInt64{Int64: payload.PostalCode, Valid: true},
	}
	address, err := ac.store.UpdateUserAddress(ctx, *addressArgs)
	if err != nil {
		tErr := custom_errors.TranslateError(err)
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(tErr.Code, tErr.Error()))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "payload": ResponseUserAddress(address)})
}

func (ac *Controller) AddUserPhone(ctx *gin.Context) {
	var payload *db.UserPhone
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(utils.INVALID_DATA, err.Error()))
		return
	}
	currentUser, _ := ctx.MustGet("User").(db.UserPayload)
	phoneArgs := &db.CreateUserPhoneParams{
		UserID:      currentUser.UserId,
		Number:      payload.Number,
		CountryCode: payload.CountryCode,
	}
	phone, err := ac.store.CreateUserPhone(ctx, *phoneArgs)
	if err != nil {
		tErr := custom_errors.TranslateError(err)
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(tErr.Code, tErr.Error()))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 201, "msg": "success", "payload": ResponseUserPhone(phone)})
}

func (ac *Controller) UpdateUserPhone(ctx *gin.Context) {
	var payload *db.UserPhone
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(utils.INVALID_DATA, err.Error()))
		return
	}
	currentUser, _ := ctx.MustGet("User").(db.UserPayload)

	phoneArgs := &db.UpdateUserPhoneParams{
		UserID:      currentUser.UserId,
		Number:      sql.NullInt64{Int64: payload.Number, Valid: true},
		CountryCode: sql.NullString{String: payload.CountryCode, Valid: true},
		OldNumber:   payload.OldNumber,
	}
	phone, err := ac.store.UpdateUserPhone(ctx, *phoneArgs)
	if err != nil {
		tErr := custom_errors.TranslateError(err)
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(tErr.Code, tErr.Error()))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 201, "msg": "success", "payload": ResponseUserPhone(phone)})
}

func (ac *Controller) UpdateAccount(ctx *gin.Context) {
	var payload *User
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(utils.INVALID_DATA, err.Error()))
		return
	}
	currentUser, _ := ctx.MustGet("User").(db.UserPayload)
	userArgs := &db.UpdateUserParams{
		ID:        currentUser.UserId,
		Username:  sql.NullString{String: payload.Username, Valid: true},
		Email:     sql.NullString{String: payload.Email, Valid: true},
		FirstName: sql.NullString{String: payload.FirstName, Valid: true},
		LastName:  sql.NullString{String: payload.LastName, Valid: true},
		Photo:     sql.NullString{String: payload.Photo, Valid: true},
	}
	_, err := ac.store.UpdateUser(ctx, *userArgs)
	if err != nil {
		tErr := custom_errors.TranslateError(err)
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ErrorResponse(tErr.Code, tErr.Error()))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
}
