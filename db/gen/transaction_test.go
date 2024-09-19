package db

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/symyzi/financial-helper/util"
)

func CreateRandomTransaction(t *testing.T, user User, category Category) Transaction {
	arg := CreateTransactionParams{
		UserID:     user.ID,
		CategoryID: category.ID,
		Amount:     util.RandomAmount(),
		Description: pgtype.Text{
			String: util.RandomString(15),
			Valid:  true,
		},
		TransactionDate: pgtype.Date{},
	}

	now := time.Now()
	arg.TransactionDate.Time = now
	arg.TransactionDate.Valid = true

	transaction, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, arg.UserID, transaction.UserID)
	require.Equal(t, arg.CategoryID, transaction.CategoryID)
	require.Equal(t, arg.Amount, transaction.Amount)
	require.Equal(t, arg.Description.String, transaction.Description.String)
	require.NotZero(t, transaction.ID)
	require.NotZero(t, transaction.CreatedAt)

	return transaction
}

func TestCreateTransaction(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)
	CreateRandomTransaction(t, user, category)
}

func TestGetTransactionByID(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)
	transaction1 := CreateRandomTransaction(t, user, category)
	transaction2, err := testQueries.GetTransactionByID(context.Background(), transaction1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transaction2)

	require.Equal(t, transaction1.ID, transaction2.ID)
	require.Equal(t, transaction1.UserID, transaction2.UserID)
	require.Equal(t, transaction1.CategoryID, transaction2.CategoryID)
	require.Equal(t, transaction1.Amount, transaction2.Amount)
	require.Equal(t, transaction1.Description.String, transaction2.Description.String)
	require.Equal(t, transaction1.TransactionDate.Time, transaction2.TransactionDate.Time)
	require.WithinDuration(t, transaction1.CreatedAt.Time, transaction2.CreatedAt.Time, time.Second)
}

func TestGetTransactionsByUserID(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)

	for i := 0; i < 5; i++ {
		CreateRandomTransaction(t, user, category)
	}

	transactions, err := testQueries.GetTransactionsByUserID(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transactions)
	require.Len(t, transactions, 5)
}

func TestGetTransactionsByCategoryID(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)

	for i := 0; i < 5; i++ {
		CreateRandomTransaction(t, user, category)
	}

	transactions, err := testQueries.GetTransactionsByCategoryID(context.Background(), category.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transactions)
	require.Len(t, transactions, 5)
}

func TestUpdateTransaction(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)
	transaction := CreateRandomTransaction(t, user, category)

	arg := UpdateTransactionParams{
		ID:              transaction.ID,
		Amount:          util.RandomAmount(),
		Description:     pgtype.Text{String: "Updated Description", Valid: true},
		CategoryID:      category.ID,
		TransactionDate: transaction.TransactionDate,
	}

	updatedTransaction, err := testQueries.UpdateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, transaction.ID, updatedTransaction.ID)
	require.Equal(t, arg.Amount, updatedTransaction.Amount)
	require.Equal(t, arg.Description.String, updatedTransaction.Description.String)
}

func TestDeleteTransaction(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)
	transaction := CreateRandomTransaction(t, user, category)

	err := testQueries.DeleteTransaction(context.Background(), transaction.ID)
	require.NoError(t, err)

	deletedTransaction, err := testQueries.GetTransactionByID(context.Background(), transaction.ID)
	require.Error(t, err)
	require.True(t, errors.Is(err, sql.ErrNoRows))
	require.Empty(t, deletedTransaction)
}

func TestGetTransactionsByDateRange(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)

	for i := 0; i < 5; i++ {
		CreateRandomTransaction(t, user, category)
	}

	now := time.Now()
	startDate := pgtype.Date{
		Time:  now.Add(-10 * 24 * time.Hour),
		Valid: true,
	}
	endDate := pgtype.Date{
		Time:  now,
		Valid: true,
	}

	transactions, err := testQueries.GetTransactionsByDateRange(context.Background(), GetTransactionsByDateRangeParams{
		UserID:            user.ID,
		TransactionDate:   startDate,
		TransactionDate_2: endDate,
	})
	require.NoError(t, err)
	require.NotEmpty(t, transactions)
	require.Len(t, transactions, 5)
}

func TestGetTotalTransactionAmountByUserID(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)

	var totalAmount int64
	for i := 0; i < 5; i++ {
		transaction := CreateRandomTransaction(t, user, category)
		totalAmount += transaction.Amount
	}

	sum, err := testQueries.GetTotalTransactionAmountByUserID(context.Background(), user.ID)
	require.NoError(t, err)
	require.Equal(t, totalAmount, sum)
}

func TestGetTransactionsWithCategoriesByUserID(t *testing.T) {
	user := CreateRandomUser(t)

	category1 := CreateRandomCategory(t)
	category2 := CreateRandomCategory(t)

	for i := 0; i < 5; i++ {
		CreateRandomTransaction(t, user, category1)
		CreateRandomTransaction(t, user, category2)
	}

	transactions, err := testQueries.GetTransactionsWithCategoriesByUserID(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transactions)
	require.Len(t, transactions, 10)
}

func TestGetTransactionsByCategoryIDAndDateRange(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)

	for i := 0; i < 5; i++ {
		CreateRandomTransaction(t, user, category)
	}

	now := time.Now()
	startDate := pgtype.Date{
		Time:  now.Add(-10 * 24 * time.Hour),
		Valid: true,
	}
	endDate := pgtype.Date{
		Time:  now,
		Valid: true,
	}

	transactions, err := testQueries.GetTransactionsByCategoryIDAndDateRange(context.Background(), GetTransactionsByCategoryIDAndDateRangeParams{
		UserID:            user.ID,
		CategoryID:        category.ID,
		TransactionDate:   startDate,
		TransactionDate_2: endDate,
	})
	require.NoError(t, err)
	require.NotEmpty(t, transactions)
	require.Len(t, transactions, 5)
}

func TestListTransactions(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)

	for i := 0; i < 15; i++ {
		CreateRandomTransaction(t, user, category)
	}

	transactions, err := testQueries.ListTransactions(context.Background(), ListTransactionsParams{
		Limit:  5,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Len(t, transactions, 5)

	transactions, err = testQueries.ListTransactions(context.Background(), ListTransactionsParams{
		Limit:  5,
		Offset: 5,
	})
	require.NoError(t, err)
	require.Len(t, transactions, 5)

	transactions, err = testQueries.ListTransactions(context.Background(), ListTransactionsParams{
		Limit:  5,
		Offset: 10,
	})
	require.NoError(t, err)
	require.Len(t, transactions, 5)
}

func TestGetTotalAmountByCategoryAndDateRange(t *testing.T) {
	user := CreateRandomUser(t)
	category1 := CreateRandomCategory(t)
	category2 := CreateRandomCategory(t)

	for i := 0; i < 5; i++ {
		CreateRandomTransaction(t, user, category1)
	}
	for i := 0; i < 3; i++ {
		CreateRandomTransaction(t, user, category2)
	}

	now := time.Now()
	startDate := pgtype.Date{
		Time:  now.Add(-10 * 24 * time.Hour),
		Valid: true,
	}
	endDate := pgtype.Date{
		Time:  now,
		Valid: true,
	}

	result, err := testQueries.GetTotalAmountByCategoryAndDateRange(context.Background(), GetTotalAmountByCategoryAndDateRangeParams{
		UserID:            user.ID,
		TransactionDate:   startDate,
		TransactionDate_2: endDate,
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Len(t, result, 2)
}
