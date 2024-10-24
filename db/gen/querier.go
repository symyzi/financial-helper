// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateBudget(ctx context.Context, arg CreateBudgetParams) (Budget, error)
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error)
	CreateExpense(ctx context.Context, arg CreateExpenseParams) (Expense, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateWallet(ctx context.Context, arg CreateWalletParams) (Wallet, error)
	DeleteBudget(ctx context.Context, id int64) error
	DeleteCategory(ctx context.Context, id int64) error
	DeleteExpense(ctx context.Context, id int64) error
	DeleteWallet(ctx context.Context, arg DeleteWalletParams) error
	GetAllCategories(ctx context.Context, owner string) ([]Category, error)
	GetBudgetByID(ctx context.Context, id int64) (Budget, error)
	GetCategoryByID(ctx context.Context, id int64) (Category, error)
	GetExpense(ctx context.Context, id int64) (Expense, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetWallet(ctx context.Context, id int64) (Wallet, error)
	ListBudgets(ctx context.Context, arg ListBudgetsParams) ([]Budget, error)
	ListExpenses(ctx context.Context, arg ListExpensesParams) ([]Expense, error)
	ListWallets(ctx context.Context, arg ListWalletsParams) ([]Wallet, error)
	UpdateBudget(ctx context.Context, arg UpdateBudgetParams) (Budget, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
	UpdateExpense(ctx context.Context, arg UpdateExpenseParams) (Expense, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
