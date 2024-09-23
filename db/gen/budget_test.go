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

func CreateRandomBudget(t *testing.T, user User, category Category) Budget {
	arg := CreateBudgetParams{
		UserID:     user.ID,
		CategoryID: category.ID,
		Amount:     util.RandomAmount(),
	}

	budget, err := testQueries.CreateBudget(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, budget)

	require.Equal(t, arg.UserID, budget.UserID)
	require.Equal(t, arg.CategoryID, budget.CategoryID)
	require.Equal(t, arg.Amount, budget.Amount)

	require.NotZero(t, budget.ID)
	require.NotZero(t, budget.CreatedAt)

	return budget
}

func TestCreateBudget(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)
	CreateRandomBudget(t, user, category)
}

func TestGetBudgetByID(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)
	budget1 := CreateRandomBudget(t, user, category)
	budget2, err := testQueries.GetBudgetByID(context.Background(), budget1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, budget2)

	require.Equal(t, budget1.ID, budget2.ID)
	require.Equal(t, budget1.UserID, budget2.UserID)
	require.Equal(t, budget1.CategoryID, budget2.CategoryID)
	require.Equal(t, budget1.Amount, budget2.Amount)
	require.WithinDuration(t, budget1.CreatedAt, budget2.CreatedAt, time.Second)
}

func TestGetBudgetByCategoryID(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)
	budget1 := CreateRandomBudget(t, user, category)
	budget2, err := testQueries.GetBudgetByCategoryID(context.Background(), budget1.CategoryID)
	require.NoError(t, err)
	require.NotEmpty(t, budget2)

	require.Equal(t, budget1.ID, budget2.ID)
	require.Equal(t, budget1.UserID, budget2.UserID)
	require.Equal(t, budget1.CategoryID, budget2.CategoryID)
	require.Equal(t, budget1.Amount, budget2.Amount)
	require.WithinDuration(t, budget1.CreatedAt, budget2.CreatedAt, time.Second)
}

func TestGetBudgetsByUserID(t *testing.T) {
	user := CreateRandomUser(t)
	categoris := make([]Category, 0)
	for i := 0; i < 10; i++ {
		categoris = append(categoris, CreateRandomCategory(t))
	}
	for i := 0; i < 10; i++ {
		CreateRandomBudget(t, user, categoris[i])
	}
	budgets, err := testQueries.GetBudgetsByUserID(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, budgets)
	require.Len(t, budgets, 10)
	for _, budget := range budgets {
		require.NotEmpty(t, budget)
	}

}

func TestUpdateBudget(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)
	budget1 := CreateRandomBudget(t, user, category)
	arg := UpdateBudgetParams{
		ID:         budget1.ID,
		CategoryID: budget1.CategoryID,
		Amount:     util.RandomAmount(),
	}
	budget2, err := testQueries.UpdateBudget(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, budget2)

	require.Equal(t, budget1.ID, budget2.ID)
	require.Equal(t, budget1.UserID, budget2.UserID)
	require.Equal(t, budget1.CategoryID, budget2.CategoryID)
	require.Equal(t, arg.Amount, budget2.Amount)
	require.WithinDuration(t, budget1.CreatedAt, budget2.CreatedAt, time.Second)
}

func TestDeleteBudget(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)
	budget1 := CreateRandomBudget(t, user, category)
	err := testQueries.DeleteBudget(context.Background(), budget1.ID)
	require.NoError(t, err)
	budget2, err := testQueries.GetBudgetByID(context.Background(), budget1.ID)
	require.Error(t, err)
	require.Empty(t, budget2)
	require.True(t, errors.Is(err, sql.ErrNoRows))
}

func TestUpdateBudgetAmount(t *testing.T) {
	user := CreateRandomUser(t)
	category := CreateRandomCategory(t)
	budget1 := CreateRandomBudget(t, user, category)

	arg := UpdateBudgetParams{
		ID:         budget1.ID,
		CategoryID: budget1.CategoryID,
		Amount:     budget1.Amount + 100,
	}

	budget2, err := testQueries.UpdateBudget(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, budget2)
	require.Equal(t, arg.Amount, budget2.Amount)
}

func TestUpdateBudgetCategory(t *testing.T) {
	user := CreateRandomUser(t)
	category1 := CreateRandomCategory(t)
	category2 := CreateRandomCategory(t)
	budget1 := CreateRandomBudget(t, user, category1)

	arg := UpdateBudgetParams{
		ID:         budget1.ID,
		CategoryID: category2.ID,
		Amount:     budget1.Amount,
	}

	budget2, err := testQueries.UpdateBudget(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, budget2)
	require.Equal(t, arg.CategoryID, budget2.CategoryID)
}
