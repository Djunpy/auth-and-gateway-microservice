package db

import (
	"auth-and-gateway-microservice/utils"
	"context"
	"fmt"
)

func (store *Store) GetUser(ctx context.Context, login string, loginType string, password string) (user User, err error) {
	switch loginType {
	case "EMAIL":
		user, err = store.GetUserByEmail(ctx, login)
		if err != nil {
			return User{}, err
		}
	case "USERNAME":
		user, err = store.GetUserByUsername(ctx, login)
		if err != nil {
			return User{}, err
		}
	}
	if err = utils.ComparePassword(user.Password, password); err != nil {
		return User{}, StoreError{Code: utils.INCORRECT_PASSWORD, Err: fmt.Errorf("incorrect password")}
	}
	return user, nil
}
