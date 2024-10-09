// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"time"
)

type Budget struct {
	ID       int64 `json:"id"`
	WalletID int64 `json:"wallet_id"`
	// must be positive
	Amount     int64     `json:"amount"`
	CategoryID int64     `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type Category struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Expense struct {
	ID       int64 `json:"id"`
	WalletID int64 `json:"wallet_id"`
	// must be positive
	Amount             int64          `json:"amount"`
	ExpenseDescription sql.NullString `json:"expense_description"`
	CategoryID         int64          `json:"category_id"`
	ExpenseDate        time.Time      `json:"expense_date"`
	CreatedAt          time.Time      `json:"created_at"`
}

type User struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	HashedPassword    string    `json:"hashed_password"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type Wallet struct {
	Name      string    `json:"name"`
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}
