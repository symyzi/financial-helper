package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/symyzi/financial-helper/util"
)

func CreateRandomBudget(t *testing.T, wallet Wallet, category Category) Budget {
	arg := CreateBudgetParams{
		WalletID:   wallet.ID,
		CategoryID: category.ID,
		Amount:     util.RandomAmount(),
	}
	budget, err := testQueries.CreateBudget(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, budget)
	require.Equal(t, arg.WalletID, budget.WalletID)
	require.Equal(t, arg.CategoryID, budget.CategoryID)
	require.Equal(t, arg.Amount, budget.Amount)
	require.NotZero(t, budget.ID)
	require.NotZero(t, budget.CreatedAt)
	return budget
}
func TestCreateBudget(t *testing.T) {
	wallet := CreateRandomWallet(t, CreateRandomUser(t))
	category := CreateRandomCategory(t)
	CreateRandomBudget(t, wallet, category)
}

func TestGetBudgetByID(t *testing.T) {
	wallet := CreateRandomWallet(t, CreateRandomUser(t))
	category := CreateRandomCategory(t)
	budget1 := CreateRandomBudget(t, wallet, category)
	budget2, err := testQueries.GetBudgetByID(context.Background(), budget1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, budget2)

	require.Equal(t, budget1.ID, budget2.ID)
	require.Equal(t, budget1.WalletID, budget2.WalletID)
	require.Equal(t, budget1.CategoryID, budget2.CategoryID)
	require.Equal(t, budget1.Amount, budget2.Amount)
	require.WithinDuration(t, budget1.CreatedAt, budget2.CreatedAt, time.Second)
}

func TestGetBudgetByCategoryID(t *testing.T) {
	wallet := CreateRandomWallet(t, CreateRandomUser(t))
	category := CreateRandomCategory(t)
	budget1 := CreateRandomBudget(t, wallet, category)
	budget2, err := testQueries.GetBudgetByCategoryID(context.Background(), budget1.CategoryID)
	require.NoError(t, err)
	require.NotEmpty(t, budget2)

	require.Equal(t, budget1.ID, budget2.ID)
	require.Equal(t, budget1.WalletID, budget2.WalletID)
	require.Equal(t, budget1.CategoryID, budget2.CategoryID)
	require.Equal(t, budget1.Amount, budget2.Amount)
	require.WithinDuration(t, budget1.CreatedAt, budget2.CreatedAt, time.Second)
}

func TestGetBudgetsByWalletID(t *testing.T) {
	wallet := CreateRandomWallet(t, CreateRandomUser(t))
	for i := 0; i < 5; i++ {
		category := CreateRandomCategory(t)
		CreateRandomBudget(t, wallet, category)
	}
	arg := wallet.ID
	budgets, err := testQueries.GetBudgetsByWalletID(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, budgets, 5)
	for _, budget := range budgets {
		require.NotEmpty(t, budget)
		require.Equal(t, wallet.ID, budget.WalletID)
	}
}

func TestUpdateBudget(t *testing.T) {
	wallet := CreateRandomWallet(t, CreateRandomUser(t))
	category := CreateRandomCategory(t)
	budget1 := CreateRandomBudget(t, wallet, category)
	arg := UpdateBudgetParams{
		ID:         budget1.ID,
		Amount:     util.RandomAmount(),
		CategoryID: category.ID,
	}
	budget2, err := testQueries.UpdateBudget(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, budget2)
	require.Equal(t, budget1.ID, budget2.ID)
	require.Equal(t, budget1.WalletID, budget2.WalletID)
	require.Equal(t, budget1.CategoryID, budget2.CategoryID)
	require.Equal(t, arg.Amount, budget2.Amount)
	require.NotEqual(t, budget1.Amount, budget2.Amount)
}

func TestDeleteBudget(t *testing.T) {
	wallet := CreateRandomWallet(t, CreateRandomUser(t))
	category := CreateRandomCategory(t)
	budget1 := CreateRandomBudget(t, wallet, category)
	err := testQueries.DeleteBudget(context.Background(), budget1.ID)
	require.NoError(t, err)
	budget2, err := testQueries.GetBudgetByID(context.Background(), budget1.ID)
	require.Error(t, err)
	require.Empty(t, budget2)
}
