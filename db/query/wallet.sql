-- name: CreateWallet :one

INSERT INTO wallets (
    name,
    owner,
    currency
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetWallet :one
SELECT * FROM wallets
WHERE id = $1 LIMIT 1;

-- name: ListWallets :many
SELECT * FROM wallets
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: DeleteWallet :exec
DELETE FROM wallets
WHERE id = $1;



