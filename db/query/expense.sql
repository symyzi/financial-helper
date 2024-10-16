-- name: CreateExpense :one
INSERT INTO expenses (
    wallet_id,
    amount,
    expense_description,
    category_id,
    expense_date
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetExpense :one
SELECT * FROM expenses
WHERE id = $1 AND wallet_id = $2 LIMIT 1;

-- name: ListExpenses :many
SELECT * FROM expenses
WHERE wallet_id = $1
ORDER BY expense_date DESC
LIMIT $2
OFFSET $3;

-- name: UpdateExpense :one
UPDATE expenses
SET amount = $2, expense_description = $3, category_id = $4, expense_date = $5
WHERE id = $1
RETURNING *;

-- name: DeleteExpense :exec
DELETE FROM expenses
WHERE id = $1;


