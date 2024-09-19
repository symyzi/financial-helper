-- name: CreateBudget :one
INSERT INTO budget (
  user_id,
  category_id,
  amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetBudgetByID :one
SELECT * FROM budget 
WHERE id = $1;

-- name: GetBudgetsByUserID :many
SELECT * FROM budget 
WHERE user_id = $1 
ORDER BY created_at DESC;

-- name: UpdateBudget :one
UPDATE budget 
SET amount = $2, category_id = $3 
WHERE id = $1
RETURNING *;

-- name: DeleteBudget :exec
DELETE FROM budget
WHERE id = $1;

-- name: GetBudgetByCategoryID :one
SELECT * FROM budget 
WHERE category_id = $1;