// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: category.sql

package db

import (
	"context"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO categories (
  name
) VALUES (
  $1
)
RETURNING id, name, created_at
`

func (q *Queries) CreateCategory(ctx context.Context, name string) (Category, error) {
	row := q.queryRow(ctx, q.createCategoryStmt, createCategory, name)
	var i Category
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories 
WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteCategoryStmt, deleteCategory, id)
	return err
}

const getAllCategories = `-- name: GetAllCategories :many
SELECT id, name, created_at FROM categories
`

func (q *Queries) GetAllCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.query(ctx, q.getAllCategoriesStmt, getAllCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCategoryByID = `-- name: GetCategoryByID :one
SELECT id, name, created_at FROM categories 
WHERE id = $1
`

func (q *Queries) GetCategoryByID(ctx context.Context, id int64) (Category, error) {
	row := q.queryRow(ctx, q.getCategoryByIDStmt, getCategoryByID, id)
	var i Category
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const updateCategory = `-- name: UpdateCategory :one
UPDATE categories 
SET name = $2 
WHERE id = $1
RETURNING id, name, created_at
`

type UpdateCategoryParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error) {
	row := q.queryRow(ctx, q.updateCategoryStmt, updateCategory, arg.ID, arg.Name)
	var i Category
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}
