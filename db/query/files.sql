-- name: CreateFile :one
INSERT INTO files(
    user_id,
    filepath
)VALUES ($1, $2)
RETURNING *;

-- name: DeleteFile :one
DELETE FROM files
WHERE user_id = $1 OR id = $2
RETURNING *;

-- name: AddFileToMessage :one
INSERT INTO message_files(
    message_id,
    file_id
)VALUES ($1, $2)
RETURNING *;
