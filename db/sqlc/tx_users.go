package db

import (
	"auth-and-gateway-microservice/utils"
	"context"
	"database/sql"
)

type UserPhone struct {
	Number      int64  `json:"number"`
	CountryCode string `json:"country_code,omitempty"`
	OldNumber   int64  `json:"old_number,omitempty"`
}

type CreateUserTxParams struct {
	Username  string `json:"username" validate:"required,min=6"`
	Email     string `json:"email" validate:"required,email"`
	Photo     string `json:"photo,omitempty"`
	Password1 string `json:"password1" validate:"required,password"`
	Password2 string `json:"password2" validate:"required,eqfield=Password1"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	UserPhone `json:"phone"`
}

func (store *Store) TxCreateUser(ctx context.Context, args *CreateUserTxParams, authSource string) error {
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		hashPass := utils.HashPassword(args.Password1)
		userArgs := &CreateUserParams{
			Username:   args.Username,
			Email:      args.Email,
			FirstName:  sql.NullString{String: args.FirstName, Valid: true},
			LastName:   sql.NullString{String: args.LastName, Valid: true},
			Password:   hashPass,
			AuthSource: authSource,
			Photo:      sql.NullString{String: args.Photo, Valid: true},
			IsActive:   sql.NullBool{Bool: true, Valid: true},
		}
		user, err := q.CreateUser(ctx, *userArgs)
		if err != nil {
			return err
		}
		phoneArgs := &CreateUserPhoneParams{
			UserID:      user.ID,
			Number:      args.Number,
			CountryCode: args.CountryCode,
		}
		_, err = q.CreateUserPhone(ctx, *phoneArgs)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
