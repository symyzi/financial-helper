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

func CreateRandomCategory(t *testing.T, user User) Category {
	arg := CreateCategoryParams{
		Name:  util.RandomString(6),
		Owner: user.Username,
	}
	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)
	require.Equal(t, arg.Name, category.Name)
	require.Equal(t, arg.Owner, category.Owner)
	require.NotZero(t, category.ID)
	require.NotZero(t, category.CreatedAt)
	return category
}

func TestCreateCategory(t *testing.T) {
	CreateRandomCategory(t, CreateRandomUser(t))
}
func TestGetCategoryByID(t *testing.T) {
	category1 := CreateRandomCategory(t, CreateRandomUser(t))
	category2, err := testQueries.GetCategoryByID(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)
	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Name, category2.Name)
	require.Equal(t, category1.Owner, category2.Owner)
	require.WithinDuration(t, category1.CreatedAt, category2.CreatedAt, time.Second)
}

func TestGetAllCategories(t *testing.T) {
	user := CreateRandomUser(t)
	for i := 0; i < 10; i++ {
		CreateRandomCategory(t, user)
	}

	categories, err := testQueries.GetAllCategories(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, categories)

	for _, category := range categories {
		require.NotEmpty(t, category)
		require.Equal(t, user.Username, category.Owner)
	}
}

func TestUpdateCategory(t *testing.T) {
	category1 := CreateRandomCategory(t, CreateRandomUser(t))
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
	category1 := CreateRandomCategory(t, CreateRandomUser(t))
	err := testQueries.DeleteCategory(context.Background(), category1.ID)
	require.NoError(t, err)
	category2, err := testQueries.GetCategoryByID(context.Background(), category1.ID)
	require.Error(t, err)
	require.Empty(t, category2)
	require.True(t, errors.Is(err, sql.ErrNoRows))
}

func TestCreateCategoryUniqueness(t *testing.T) {
	categoryName := "Test Category" + util.RandomString(6)
	user := CreateRandomUser(t)
	arg1 := CreateCategoryParams{
		Name:  categoryName,
		Owner: user.Username,
	}
	arg2 := CreateCategoryParams{
		Name:  categoryName,
		Owner: user.Username,
	}
	category1, err := testQueries.CreateCategory(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, category1)

	category2, err := testQueries.CreateCategory(context.Background(), arg2)
	require.Error(t, err) // Expecting an error due to unique constraint violation
	require.Empty(t, category2)
}
