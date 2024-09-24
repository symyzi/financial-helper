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

func CreateRandomCategory(t *testing.T) Category {
	arg := util.RandomString(6)
	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)
	require.Equal(t, arg, category.Name)
	require.NotZero(t, category.ID)
	require.NotZero(t, category.CreatedAt)
	return category
}

func TestCreateCategory(t *testing.T) {
	CreateRandomCategory(t)
}

func TestGetCategoryByID(t *testing.T) {
	category1 := CreateRandomCategory(t)
	category2, err := testQueries.GetCategoryByID(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)
	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Name, category2.Name)
	require.WithinDuration(t, category1.CreatedAt, category2.CreatedAt, time.Second)
}

func TestGetAllCategories(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomCategory(t)
	}

	categories, err := testQueries.GetAllCategories(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, categories)

	for _, category := range categories {
		require.NotEmpty(t, category)
	}
}

func TestUpdateCategory(t *testing.T) {
	category1 := CreateRandomCategory(t)
	arg := UpdateCategoryParams{
		ID:   category1.ID,
		Name: util.RandomString(6),
	}
	category2, err := testQueries.UpdateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)
	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, arg.Name, category2.Name)
	require.WithinDuration(t, category1.CreatedAt, category2.CreatedAt, time.Second)
}

func TestDeleteCategory(t *testing.T) {
	category1 := CreateRandomCategory(t)
	err := testQueries.DeleteCategory(context.Background(), category1.ID)
	require.NoError(t, err)
	category2, err := testQueries.GetCategoryByID(context.Background(), category1.ID)
	require.Error(t, err)
	require.Empty(t, category2)
	require.True(t, errors.Is(err, sql.ErrNoRows))
}

func TestCreateCategoryUniqueness(t *testing.T) {
	// Create two categories with the same name
	categoryName := "Test Category" + util.RandomString(6)
	category1, err := testQueries.CreateCategory(context.Background(), categoryName)
	require.NoError(t, err)
	require.NotEmpty(t, category1)

	category2, err := testQueries.CreateCategory(context.Background(), categoryName)
	require.Error(t, err) // Expecting an error due to unique constraint violation
	require.Empty(t, category2)
}
