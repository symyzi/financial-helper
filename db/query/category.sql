-- name: CreateCategory :one
INSERT INTO categories (
  name,
  owner
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetCategoryByID :one
SELECT * FROM categories 
WHERE id = $1;

-- name: GetAllCategories :many
SELECT * FROM categories
WHERE owner = $1;


-- name: UpdateCategory :one
UPDATE categories 
SET name = $2 
WHERE id = $1
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories 
WHERE id = $1;