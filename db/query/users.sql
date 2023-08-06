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

-- name: UpdateUserPhone :one
UPDATE phones
SET
    number = COALESCE(sqlc.narg('number'), number),
    country_code = COALESCE(sqlc.narg('country_code'), country_code)
WHERE user_id = sqlc.arg('user_id') OR number = sqlc.arg('old_number')
RETURNING *;

-- name: CreateUserAddress :one
INSERT INTO address(
    user_id,
    city,
    street,
    postal_code
)VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: UpdateUserAddress :one
UPDATE address
SET
    city = COALESCE(sqlc.narg('city'), city),
    street = COALESCE(sqlc.narg('street'), street),
    postal_code = COALESCE(sqlc.narg('postal_code'), postal_code)
WHERE user_id = sqlc.arg('user_id')
RETURNING *;