-- name: CreateChat :one
INSERT INTO chats (is_deleted) VALUES (false)
RETURNING *;

-- name: GetChatByUsers :one
SELECT c.*
FROM chats c
         JOIN chat_participants p ON c.id = p.chat_id
WHERE p.user_id IN ($1, $2)
  AND c.is_deleted = false;

-- name: AddUserToChat :one
INSERT INTO chat_participants (
    user_id,
    chat_id
)VALUES ($1, $2) RETURNING *;

-- name: DeleteChatFromUser :one
UPDATE chat_participants
    SET
        is_deleted = true
WHERE user_id = $1 OR chat_id = $2
RETURNING *;

-- name: CreateMessage :one
INSERT INTO messages (
    chat_id,
    sender_id,
    text_message
) VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateMessage :one
UPDATE messages
SET
    text_message = COALESCE(sqlc.narg('text_message'), text_message),
    is_deleted = COALESCE(sqlc.narg('is_deleted'), is_deleted)
WHERE sender_id = sqlc.arg('sender_id') OR chat_id = sqlc.arg('chat_id')
RETURNING *;

-- name: CreateMessageStatus :one
INSERT INTO message_status(
    message_id,
    is_read,
    is_delivered
)VALUES($1, $2, $3)
RETURNING *;
