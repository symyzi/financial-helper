-- name: CreateTransaction :one
INSERT INTO transactions (
  user_id,
  amount,
  description,
  category_id,
  transaction_date
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetTransactionByID :one
SELECT * FROM transactions
WHERE id = $1;

-- name: ListTransactions :many
SELECT * FROM transactions
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetTransactionsByUserID :many
SELECT * FROM transactions 
WHERE user_id = $1 
ORDER BY transaction_date DESC;

-- name: UpdateTransaction :one
UPDATE transactions 
SET amount = $2, description = $3, category_id = $4, transaction_date = $5 
WHERE id = $1
RETURNING *;

-- name: DeleteTransaction :exec
DELETE FROM transactions 
WHERE id = $1;

-- name: GetTransactionsByDateRange :many
SELECT * FROM transactions
WHERE user_id = $1
AND transaction_date BETWEEN $2 AND $3
ORDER BY transaction_date DESC;

-- name: GetTotalTransactionAmountByUserID :one
SELECT SUM(amount) 
FROM transactions 
WHERE user_id = $1;

-- name: GetTransactionsByCategoryID :many
SELECT id, user_id, amount, description, transaction_date, created_at 
FROM transactions 
WHERE category_id = $1 
ORDER BY transaction_date DESC;

-- name: GetTotalAmountByCategoryAndDateRange :many
SELECT category_id, SUM(amount) AS total_amount 
FROM transactions 
WHERE user_id = $1 AND transaction_date BETWEEN $2 AND $3 
GROUP BY category_id;

-- name: GetTransactionsWithCategoriesByUserID :many
SELECT t.id, t.user_id, t.amount, t.description, 
       t.transaction_date, c.name AS category_name 
FROM transactions t
JOIN categories c ON t.category_id = c.id 
WHERE t.user_id = $1 
ORDER BY t.transaction_date DESC;

-- name: GetTransactionsByCategoryIDAndDateRange :many
SELECT * FROM transactions 
WHERE user_id = $1
AND category_id = $2
AND transaction_date BETWEEN $3 AND $4
ORDER BY transaction_date DESC;

-- name: GetMonthlyTransactionStatistics :many
SELECT DATE_TRUNC('month', transaction_date) AT TIME ZONE 'UTC' AS month, 
       SUM(amount) AS total_amount 
FROM transactions 
WHERE user_id = $1 
GROUP BY month 
ORDER BY month DESC;


