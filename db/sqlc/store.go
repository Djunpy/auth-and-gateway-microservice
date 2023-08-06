package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db, Queries: New(db)}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil && !errors.Is(rbErr, sql.ErrTxDone) {
			return fmt.Errorf("failed to rollback transaction: %v", rbErr)
		}
		return err
	}
	return tx.Commit()
}

type UserPayload struct {
	UserId     int32     `json:"user_id"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	IsSuper    bool      `json:"is_super"`
	IsActive   bool      `json:"is_active"`
	IsStaff    bool      `json:"is_staff"`
	IsDeleted  bool      `json:"is_deleted"`
	AuthSource string    `json:"auth_source"`
	DateJoined time.Time `json:"date_joined"`
}

func UserPayloadToCookie(user User) UserPayload {
	return UserPayload{
		UserId:     user.ID,
		Email:      user.Email,
		Username:   user.Username,
		IsSuper:    user.IsSuperuser.Bool,
		IsActive:   user.IsActive.Bool,
		IsStaff:    user.IsStaff.Bool,
		IsDeleted:  user.IsDeleted.Bool,
		AuthSource: user.AuthSource,
		DateJoined: user.DateJoined.Time,
	}
}
