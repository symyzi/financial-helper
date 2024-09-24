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

func CreateRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: util.RandomUsername(),
		Email:    util.RandomEmail(),
		Password: util.RandomPassword(),
		Currency: util.RandomCurrency(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Currency, user.Currency)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.ID)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUserByID(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUserByID(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Currency, user2.Currency)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestGetUserByUsername(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUserByUsername(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Currency, user2.Currency)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	arg := UpdateUserParams{
		ID:       user1.ID,
		Username: util.RandomUsername(),
		Email:    user1.Email,
		Password: user1.Password,
		Currency: user1.Currency,
	}
	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, arg.Username, user2.Username)
	require.Equal(t, arg.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Currency, user2.Currency)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestDeleteUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUserByID(context.Background(), user1.ID)
	require.Error(t, err)
	require.True(t, errors.Is(err, sql.ErrNoRows))
	require.Empty(t, user2)
}

func TestCheckUserCredentials(t *testing.T) {
	hashedPassword := util.RandomString(6)

	user, err := testQueries.CreateUser(context.Background(), CreateUserParams{
		Username: util.RandomUsername(),
		Email:    util.RandomEmail(),
		Password: hashedPassword,
		Currency: util.RandomCurrency(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, user)

	user2, err := testQueries.CheckUserCredentials(context.Background(), CheckUserCredentialsParams{
		Username: user.Username,
		Password: hashedPassword,
	})
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user.ID, user2.ID)
	require.Equal(t, user.Username, user2.Username)
	require.Equal(t, user.Email, user2.Email)

	wrongPassword := util.RandomString(6)
	_, err = testQueries.CheckUserCredentials(context.Background(), CheckUserCredentialsParams{
		Username: user.Username,
		Password: wrongPassword,
	})
	require.Error(t, err)
}

func TestUpdateUserPassword(t *testing.T) {
	createdUser := CreateRandomUser(t)

	arg := UpdateUserPasswordParams{
		ID:       createdUser.ID,
		Password: util.RandomString(6),
	}

	user, err := testQueries.UpdateUserPassword(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, createdUser.ID, user.ID)
	require.Equal(t, createdUser.Username, user.Username)
	require.Equal(t, createdUser.Email, user.Email)
	require.WithinDuration(t, createdUser.CreatedAt, user.CreatedAt, time.Second)
}
