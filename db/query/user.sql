-- name: CreateUser :one
INSERT INTO users (
  username,
  email,
  password,
  currency
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users 
WHERE id = $1;

-- name: GetUserByUsername :one
SELECT * FROM users 
WHERE username = $1;

-- name: UpdateUser :one
UPDATE users 
SET username = $2, email = $3, password = $4, currency = $5 
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = $1;

-- name: UpdateUserPassword :one
UPDATE users 
SET password = $2 
WHERE id = $1 
RETURNING id, username, email, created_at, currency;

-- name: CheckUserCredentials :one
SELECT id, username, email 
FROM users 
WHERE username = $1 AND password = $2;