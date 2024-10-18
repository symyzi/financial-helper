package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/symyzi/financial-helper/util"
)

func CreateRandomExpense(t *testing.T, wallet Wallet, category Category) Expense {
	arg := CreateExpenseParams{
		WalletID:           wallet.ID,
		Amount:             util.RandomAmount(),
		ExpenseDescription: util.RandomString(12),
		CategoryID:         category.ID,
	}
	expense, err := testQueries.CreateExpense(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, expense)
	require.Equal(t, arg.WalletID, expense.WalletID)
	require.Equal(t, arg.Amount, expense.Amount)
	require.Equal(t, arg.ExpenseDescription, expense.ExpenseDescription)
	require.Equal(t, arg.CategoryID, expense.CategoryID)
	require.NotZero(t, expense.ID)
	require.NotZero(t, expense.CreatedAt)
	return expense
}
func TestCreateExpense(t *testing.T) {
	wallet := CreateRandomWallet(t, CreateRandomUser(t))
	category := CreateRandomCategory(t)
	CreateRandomExpense(t, wallet, category)
}

func TestGetExpenseByID(t *testing.T) {
	wallet := CreateRandomWallet(t, CreateRandomUser(t))
	category := CreateRandomCategory(t)
	expense1 := CreateRandomExpense(t, wallet, category)
	expense2, err := testQueries.GetExpense(context.Background(), expense1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, expense2)
	require.Equal(t, expense1.ID, expense2.ID)
	require.Equal(t, expense1.WalletID, expense2.WalletID)
	require.Equal(t, expense1.Amount, expense2.Amount)
	require.Equal(t, expense1.ExpenseDescription, expense2.ExpenseDescription)
	require.Equal(t, expense1.CategoryID, expense2.CategoryID)
	require.WithinDuration(t, expense1.CreatedAt, expense2.CreatedAt, time.Second)
}

func TestDeleteExpense(t *testing.T) {
	wallet := CreateRandomWallet(t, CreateRandomUser(t))
	category := CreateRandomCategory(t)
	expense1 := CreateRandomExpense(t, wallet, category)
	err := testQueries.DeleteExpense(context.Background(), expense1.ID)
	require.NoError(t, err)
	expense2, err := testQueries.GetExpense(context.Background(), expense1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, expense2)
}

// func TestGetExpensesByWalletID(t *testing.T) {
// 	wallet := CreateRandomWallet(t, CreateRandomUser(t))
// 	for i := 0; i < 5; i++ {
// 		category := CreateRandomCategory(t)
// 		CreateRandomExpense(t, wallet, category)
// 	}
// 	arg := wallet.ID
// 	expenses, err := testQueries.GetExpensesByWalletID(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.Len(t, expenses, 5)
// 	for _, expense := range expenses {
// 		require.NotEmpty(t, expense)
// 		require.Equal(t, wallet.ID, expense.WalletID)
// 	}
// }

func TestUpdateExpense(t *testing.T) {
	wallet := CreateRandomWallet(t, CreateRandomUser(t))
	category := CreateRandomCategory(t)
	expense1 := CreateRandomExpense(t, wallet, category)
	arg := UpdateExpenseParams{
		ID:                 expense1.ID,
		Amount:             util.RandomAmount(),
		ExpenseDescription: util.RandomString(12),
		CategoryID:         category.ID,
	}
	expense2, err := testQueries.UpdateExpense(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, expense2)
	require.Equal(t, expense1.ID, expense2.ID)
	require.Equal(t, expense1.WalletID, expense2.WalletID)
	require.Equal(t, arg.Amount, expense2.Amount)
	require.Equal(t, arg.ExpenseDescription, expense2.ExpenseDescription)
	require.Equal(t, arg.CategoryID, expense2.CategoryID)
	require.NotEqual(t, expense1.Amount, expense2.Amount)
	require.NotEqual(t, expense1.ExpenseDescription, expense2.ExpenseDescription)
}

func TestListExpenses(t *testing.T) {
	wallet := CreateRandomWallet(t, CreateRandomUser(t))
	for i := 0; i < 10; i++ {
		category := CreateRandomCategory(t)
		CreateRandomExpense(t, wallet, category)
	}
	arg := ListExpensesParams{
		Limit:  5,
		Offset: 5,
	}

	expenses, err := testQueries.ListExpenses(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, expenses, 5)
	for _, expense := range expenses {
		require.NotEmpty(t, expense)
		require.Equal(t, wallet.ID, expense.WalletID)
	}
}
