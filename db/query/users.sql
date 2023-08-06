-- name: CreateUser :one
INSERT INTO users (
        username,
        email,
        photo,
        password,
        last_name,
        first_name,
        auth_source,
        is_active
) VALUES
      ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;


-- name: UpdateUser :one
UPDATE users
SET
    username = COALESCE(sqlc.narg('username'), username),
    email = COALESCE(sqlc.narg('email'), email),
    photo = COALESCE(sqlc.narg('photo'), photo),
    first_name = COALESCE(sqlc.narg('first_name'), first_name),
    last_name = COALESCE(sqlc.narg('last_name'), last_name)
WHERE id = sqlc.arg('id')
RETURNING *;


-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1;

-- name: CreateUserPhone :one
INSERT INTO phones (
    user_id,
    number,
    country_code
) VALUES ($1, $2, $3)
RETURNING *;

