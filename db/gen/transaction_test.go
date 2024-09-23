package db

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/symyzi/financial-helper/util"
)

func CreateRandomTransaction(t *testing.T, user User, category Category) Transaction {
	arg := CreateTransactionParams{
		UserID:          user.ID,
		Amount:          util.RandomAmount(),
		Description:     sql.NullString{String: util.RandomString(12), Valid: true},
		CategoryID:      category.ID,
		TransactionDate: time.Now(),
	}

	transaction, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, arg.UserID, transaction.UserID)
	require.Equal(t, arg.Amount, transaction.Amount)
	require.Equal(t, arg.Description, transaction.Description)
	require.Equal(t, arg.CategoryID, transaction.CategoryID)
	require.WithinDuration(t, arg.TransactionDate, transaction.TransactionDate, time.Second)

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
	require.Equal(t, transaction1.Amount, transaction2.Amount)
	require.Equal(t, transaction1.Description, transaction2.Description)
	require.Equal(t, transaction1.CategoryID, transaction2.CategoryID)
	require.WithinDuration(t, transaction1.TransactionDate, transaction2.TransactionDate, time.Second)
	require.WithinDuration(t, transaction1.CreatedAt, transaction2.CreatedAt, time.Second)
}

func TestListTransactions(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)
	for i := 0; i < 10; i++ {
		CreateRandomTransaction(t, user, category)
	}

	arg := ListTransactionsParams{
		Limit:  5,
		Offset: 5,
	}

	transactions, err := testQueries.ListTransactions(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transactions, 5)

	for _, transaction := range transactions {
		require.NotEmpty(t, transaction)
	}
}

func TestUpdateTransaction(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)
	transaction1 := CreateRandomTransaction(t, user, category)

	arg := UpdateTransactionParams{
		ID:              transaction1.ID,
		Amount:          util.RandomAmount(),
		Description:     sql.NullString{String: util.RandomString(12), Valid: true},
		CategoryID:      category.ID,
		TransactionDate: time.Now(),
	}

	transaction2, err := testQueries.UpdateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction2)

	require.Equal(t, arg.ID, transaction2.ID)
	require.Equal(t, arg.Amount, transaction2.Amount)
	require.Equal(t, arg.Description, transaction2.Description)
	require.Equal(t, arg.CategoryID, transaction2.CategoryID)
	require.WithinDuration(t, arg.TransactionDate, transaction2.TransactionDate, time.Second)
	require.WithinDuration(t, transaction1.CreatedAt, transaction2.CreatedAt, time.Second)
}

func TestDeleteTransaction(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)
	transaction1 := CreateRandomTransaction(t, user, category)
	err := testQueries.DeleteTransaction(context.Background(), transaction1.ID)
	require.NoError(t, err)

	transaction2, err := testQueries.GetTransactionByID(context.Background(), transaction1.ID)
	require.Error(t, err)
	require.Empty(t, transaction2)
	require.True(t, errors.Is(err, sql.ErrNoRows))
}

func TestGetTransactionsByUserID(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)
	for i := 0; i < 10; i++ {
		CreateRandomTransaction(t, user, category)
	}

	transactions, err := testQueries.GetTransactionsByUserID(context.Background(), user.ID)
	require.NoError(t, err)
	require.Len(t, transactions, 10)

	for _, transaction := range transactions {
		require.NotEmpty(t, transaction)
	}
}

func TestGetTransactionsByCategoryID(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)
	for i := 0; i < 10; i++ {
		CreateRandomTransaction(t, user, category)
	}

	transactions, err := testQueries.GetTransactionsByCategoryID(context.Background(), category.ID)
	require.NoError(t, err)
	require.Len(t, transactions, 10)

	for _, transaction := range transactions {
		require.NotEmpty(t, transaction)
	}
}

func TestGetTransactionsByDateRange(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)

	// Create transactions within a specific date range
	startDate := time.Now().AddDate(0, -1, 0) // Start date a month ago
	endDate := time.Now()
	for i := 0; i < 5; i++ {
		transactionDate := startDate.AddDate(0, 0, i) // Add a day to each transaction
		transaction := CreateTransactionParams{
			UserID:          user.ID,
			Amount:          util.RandomAmount(),
			Description:     sql.NullString{String: util.RandomString(12), Valid: true},
			CategoryID:      category.ID,
			TransactionDate: transactionDate,
		}
		_, err := testQueries.CreateTransaction(context.Background(), transaction)
		require.NoError(t, err)
	}

	// Retrieve transactions within the date range
	transactions, err := testQueries.GetTransactionsByDateRange(context.Background(), GetTransactionsByDateRangeParams{
		UserID:            user.ID,
		TransactionDate:   startDate,
		TransactionDate_2: endDate,
	})
	require.NoError(t, err)
	require.Len(t, transactions, 5)

	for _, transaction := range transactions {
		require.NotEmpty(t, transaction)
		require.True(t, transaction.TransactionDate.After(startDate) || transaction.TransactionDate.Equal(startDate))
		require.True(t, transaction.TransactionDate.Before(endDate) || transaction.TransactionDate.Equal(endDate))
	}
}

func TestGetTransactionsByCategoryIDAndDateRange(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)

	// Create transactions within a specific date range
	startDate := time.Now().AddDate(0, -1, 0) // Start date a month ago
	endDate := time.Now()
	for i := 0; i < 5; i++ {
		transactionDate := startDate.AddDate(0, 0, i) // Add a day to each transaction
		transaction := CreateTransactionParams{
			UserID:          user.ID,
			Amount:          util.RandomAmount(),
			Description:     sql.NullString{String: util.RandomString(12), Valid: true},
			CategoryID:      category.ID,
			TransactionDate: transactionDate,
		}
		_, err := testQueries.CreateTransaction(context.Background(), transaction)
		require.NoError(t, err)
	}

	// Retrieve transactions within the date range
	transactions, err := testQueries.GetTransactionsByCategoryIDAndDateRange(context.Background(), GetTransactionsByCategoryIDAndDateRangeParams{
		UserID:            user.ID,
		CategoryID:        category.ID,
		TransactionDate:   startDate,
		TransactionDate_2: endDate,
	})
	require.NoError(t, err)
	require.Len(t, transactions, 5)

	for _, transaction := range transactions {
		require.NotEmpty(t, transaction)
		require.True(t, transaction.TransactionDate.After(startDate) || transaction.TransactionDate.Equal(startDate))
		require.True(t, transaction.TransactionDate.Before(endDate) || transaction.TransactionDate.Equal(endDate))
	}
}

func TestGetTotalTransactionAmountByUserID(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)
	var expectedTotalAmount int64 = 0
	for i := 0; i < 10; i++ {
		transaction := CreateRandomTransaction(t, user, category)
		expectedTotalAmount += transaction.Amount
	}

	totalAmount, err := testQueries.GetTotalTransactionAmountByUserID(context.Background(), user.ID)
	require.NoError(t, err)
	require.Equal(t, expectedTotalAmount, totalAmount)
}
