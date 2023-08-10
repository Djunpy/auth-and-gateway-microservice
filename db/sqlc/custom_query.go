package db

import (
	"auth-and-gateway-microservice/utils"
	"context"
	"database/sql"
	"errors"
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

func (store *Store) GetOrCreateChat(ctx context.Context, participantsIds []int32) (chat Chat, err error) {
	getChatArgs := &GetChatByUsersParams{
		UserID:   sql.NullInt32{Int32: participantsIds[0], Valid: true},
		UserID_2: sql.NullInt32{Int32: participantsIds[1], Valid: true},
	}
	chat, err = store.GetChatByUsers(ctx, *getChatArgs)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			chat, err = store.CreateChat(ctx)
			if err != nil {
				return chat, err
			}
			for _, id := range participantsIds {
				participantArgs := &AddUserToChatParams{
					UserID: sql.NullInt32{Int32: id, Valid: true},
					ChatID: chat.ID,
				}
				_, err = store.AddUserToChat(ctx, *participantArgs)
				if err != nil {
					return chat, err
				}
			}
		} else {
			return chat, err
		}
	}
	return chat, err
}
