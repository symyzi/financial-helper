// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package gen

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Budget struct {
	ID     int64
	UserID int64
	// must be positive
	Amount     int64
	CategoryID int64
	CreatedAt  pgtype.Timestamptz
}

type Category struct {
	ID        int64
	Name      string
	CreatedAt pgtype.Timestamptz
}

type Transaction struct {
	ID     int64
	UserID int64
	// must be positive
	Amount          int64
	Description     pgtype.Text
	CategoryID      int64
	TransactionDate pgtype.Date
	CreatedAt       pgtype.Timestamptz
}

type User struct {
	ID        int64
	Username  string
	Email     string
	Password  string
	CreatedAt pgtype.Timestamptz
	Currency  string
}
