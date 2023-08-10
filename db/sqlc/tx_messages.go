package db

import (
	"context"
	"database/sql"
)

type CustomMessage struct {
	RecipientId int32    `json:"recipient_id" binding:"required"`
	TextMessage string   `json:"text" binding:"required"`
	File        []string `json:"file,omitempty"`
}

func (store *Store) txCreateMessage(ctx context.Context, args CustomMessage, chatId, senderId int32) error {
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		messageArgs := &CreateMessageParams{
			ChatID:      chatId,
			SenderID:    sql.NullInt32{Int32: senderId, Valid: true},
			TextMessage: args.TextMessage,
		}
		message, err := q.CreateMessage(ctx, *messageArgs)
		if err != nil {
			return err
		}
		statusArgs := &CreateMessageStatusParams{
			MessageID:   message.ID,
			IsDelivered: sql.NullBool{Bool: true, Valid: true},
		}
		_, err = q.CreateMessageStatus(ctx, *statusArgs)
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
