// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: expense.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createExpense = `-- name: CreateExpense :one
INSERT INTO expenses (
    wallet_id,
    amount,
    expense_description,
    category_id,
    expense_date
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, wallet_id, amount, expense_description, category_id, expense_date, created_at
`

type CreateExpenseParams struct {
	WalletID           int64          `json:"wallet_id"`
	Amount             int64          `json:"amount"`
	ExpenseDescription sql.NullString `json:"expense_description"`
	CategoryID         int64          `json:"category_id"`
	ExpenseDate        time.Time      `json:"expense_date"`
}

func (q *Queries) CreateExpense(ctx context.Context, arg CreateExpenseParams) (Expense, error) {
	row := q.queryRow(ctx, q.createExpenseStmt, createExpense,
		arg.WalletID,
		arg.Amount,
		arg.ExpenseDescription,
		arg.CategoryID,
		arg.ExpenseDate,
	)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.WalletID,
		&i.Amount,
		&i.ExpenseDescription,
		&i.CategoryID,
		&i.ExpenseDate,
		&i.CreatedAt,
	)
	return i, err
}

const deleteExpense = `-- name: DeleteExpense :exec
DELETE FROM expenses
WHERE id = $1
`

func (q *Queries) DeleteExpense(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteExpenseStmt, deleteExpense, id)
	return err
}

const getExpenseByID = `-- name: GetExpenseByID :one
SELECT id, wallet_id, amount, expense_description, category_id, expense_date, created_at FROM expenses
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetExpenseByID(ctx context.Context, id int64) (Expense, error) {
	row := q.queryRow(ctx, q.getExpenseByIDStmt, getExpenseByID, id)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.WalletID,
		&i.Amount,
		&i.ExpenseDescription,
		&i.CategoryID,
		&i.ExpenseDate,
		&i.CreatedAt,
	)
	return i, err
}

const listExpenses = `-- name: ListExpenses :many
SELECT id, wallet_id, amount, expense_description, category_id, expense_date, created_at FROM expenses
WHERE wallet_id = $1
ORDER BY expense_date DESC
LIMIT $2
OFFSET $3
`

type ListExpensesParams struct {
	WalletID int64 `json:"wallet_id"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *Queries) ListExpenses(ctx context.Context, arg ListExpensesParams) ([]Expense, error) {
	rows, err := q.query(ctx, q.listExpensesStmt, listExpenses, arg.WalletID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Expense{}
	for rows.Next() {
		var i Expense
		if err := rows.Scan(
			&i.ID,
			&i.WalletID,
			&i.Amount,
			&i.ExpenseDescription,
			&i.CategoryID,
			&i.ExpenseDate,
			&i.CreatedAt,
		); err != nil {
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

const updateExpense = `-- name: UpdateExpense :one
UPDATE expenses
SET amount = $2, expense_description = $3, category_id = $4, expense_date = $5
WHERE id = $1
RETURNING id, wallet_id, amount, expense_description, category_id, expense_date, created_at
`

type UpdateExpenseParams struct {
	ID                 int64          `json:"id"`
	Amount             int64          `json:"amount"`
	ExpenseDescription sql.NullString `json:"expense_description"`
	CategoryID         int64          `json:"category_id"`
	ExpenseDate        time.Time      `json:"expense_date"`
}

func (q *Queries) UpdateExpense(ctx context.Context, arg UpdateExpenseParams) (Expense, error) {
	row := q.queryRow(ctx, q.updateExpenseStmt, updateExpense,
		arg.ID,
		arg.Amount,
		arg.ExpenseDescription,
		arg.CategoryID,
		arg.ExpenseDate,
	)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.WalletID,
		&i.Amount,
		&i.ExpenseDescription,
		&i.CategoryID,
		&i.ExpenseDate,
		&i.CreatedAt,
	)
	return i, err
}
