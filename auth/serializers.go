package auth

import (
	db "auth-and-gateway-microservice/db/sqlc"
	"time"
)

type SignInInputSerializer struct {
	Email    string `json:"email,omitempty" binding:"email"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username,omitempty"`
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

func UserPayloadToCookie(user db.User) UserPayload {
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
