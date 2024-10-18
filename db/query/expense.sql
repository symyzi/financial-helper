-- name: CreateExpense :one
INSERT INTO expenses (
    wallet_id,
    amount,
    expense_description,
    category_id
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetExpense :one
SELECT * FROM expenses
WHERE id = $1 LIMIT 1;

-- name: ListExpenses :many
SELECT * FROM expenses
LIMIT $1
OFFSET $2;

-- name: UpdateExpense :one
UPDATE expenses
SET amount = $2, expense_description = $3, category_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteExpense :exec
DELETE FROM expenses
WHERE id = $1;


