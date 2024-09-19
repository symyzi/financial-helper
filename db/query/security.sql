-- name: UpdateUserPassword :one
UPDATE users 
SET password = $2 
WHERE id = $1 
RETURNING id, username, email, created_at, currency;

-- name: CheckUserCredentials :one
SELECT id, username, email 
FROM users 
WHERE username = $1 AND password = $2;