// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"database/sql"
)

type Address struct {
	ID         int32          `json:"id"`
	UserID     int32          `json:"user_id"`
	City       string         `json:"city"`
	Street     sql.NullString `json:"street"`
	PostalCode sql.NullInt64  `json:"postal_code"`
}

type Chat struct {
	ID        int32        `json:"id"`
	IsDeleted sql.NullBool `json:"is_deleted"`
	CreateAt  sql.NullTime `json:"create_at"`
}

type ChatParticipant struct {
	ID        int32         `json:"id"`
	UserID    sql.NullInt32 `json:"user_id"`
	ChatID    int32         `json:"chat_id"`
	IsDeleted sql.NullBool  `json:"is_deleted"`
}

type File struct {
	ID       int32          `json:"id"`
	UserID   int32          `json:"user_id"`
	Filename sql.NullString `json:"filename"`
	Filepath string         `json:"filepath"`
}

type Group struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	ID          int32         `json:"id"`
	ChatID      int32         `json:"chat_id"`
	SenderID    sql.NullInt32 `json:"sender_id"`
	TextMessage string        `json:"text_message"`
	IsDeleted   sql.NullBool  `json:"is_deleted"`
	SentAt      sql.NullTime  `json:"sent_at"`
}

type MessageFile struct {
	ID        int32 `json:"id"`
	MessageID int32 `json:"message_id"`
	FileID    int32 `json:"file_id"`
}

type MessageStatus struct {
	ID          int32        `json:"id"`
	MessageID   int32        `json:"message_id"`
	IsRead      sql.NullBool `json:"is_read"`
	IsDelivered sql.NullBool `json:"is_delivered"`
}

type Permission struct {
	ID       int32          `json:"id"`
	Name     string         `json:"name"`
	Codename sql.NullString `json:"codename"`
}

type Phone struct {
	ID          int32        `json:"id"`
	UserID      int32        `json:"user_id"`
	Number      int64        `json:"number"`
	CountryCode string       `json:"country_code"`
	Verified    sql.NullBool `json:"verified"`
	CreateAt    sql.NullTime `json:"create_at"`
}

type ReplyToRating struct {
	ID        int32        `json:"id"`
	UserID    int32        `json:"user_id"`
	RatingID  int32        `json:"rating_id"`
	ReplyText string       `json:"reply_text"`
	CreateAt  sql.NullTime `json:"create_at"`
}

type User struct {
	ID            int32          `json:"id"`
	Username      string         `json:"username"`
	Email         string         `json:"email"`
	VerifiedEmail sql.NullBool   `json:"verified_email"`
	Photo         sql.NullString `json:"photo"`
	Password      string         `json:"password"`
	FirstName     sql.NullString `json:"first_name"`
	LastName      sql.NullString `json:"last_name"`
	IsStaff       sql.NullBool   `json:"is_staff"`
	IsActive      sql.NullBool   `json:"is_active"`
	IsSuperuser   sql.NullBool   `json:"is_superuser"`
	IsDeleted     sql.NullBool   `json:"is_deleted"`
	AuthSource    string         `json:"auth_source"`
	UpdateAt      sql.NullTime   `json:"update_at"`
	DateJoined    sql.NullTime   `json:"date_joined"`
}

type UserGroup struct {
	ID       int32        `json:"id"`
	UserID   int32        `json:"user_id"`
	GroupID  int32        `json:"group_id"`
	CreateAt sql.NullTime `json:"create_at"`
}

type UserPermission struct {
	ID           int32        `json:"id"`
	UserID       int32        `json:"user_id"`
	PermissionID int32        `json:"permission_id"`
	CreateAt     sql.NullTime `json:"create_at"`
}

type UserRating struct {
	ID          int32         `json:"id"`
	UserID      int32         `json:"user_id"`
	RaterID     sql.NullInt32 `json:"rater_id"`
	RatingValue int32         `json:"rating_value"`
	Comment     string        `json:"comment"`
	IsDeleted   sql.NullBool  `json:"is_deleted"`
	CreateAt    sql.NullTime  `json:"create_at"`
}
