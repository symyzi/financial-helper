-- name: CreateBudget :one
INSERT INTO budgets (
  wallet_id,
  category_id,
  amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetBudgetByID :one
SELECT * FROM budgets 
WHERE id = $1;

-- name: GetBudgetsByWalletID :many
SELECT * FROM budgets 
WHERE wallet_id = $1 
ORDER BY created_at DESC;

-- name: UpdateBudget :one
UPDATE budgets 
SET amount = $2, category_id = $3 
WHERE id = $1
RETURNING *;

-- name: DeleteBudget :exec
DELETE FROM budgets
WHERE id = $1;

-- name: GetBudgetByCategoryID :one
SELECT * FROM budgets 
WHERE category_id = $1;