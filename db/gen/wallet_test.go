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

func CreateRandomWallet(t *testing.T, user User) Wallet {
	arg := CreateWalletParams{
		Name:     util.RandomString(6),
		Owner:    user.Username,
		Currency: util.RandomCurrency(),
	}

	wallet, err := testQueries.CreateWallet(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, wallet)

	require.Equal(t, arg.Owner, wallet.Owner)
	require.Equal(t, arg.Name, wallet.Name)
	require.Equal(t, arg.Currency, wallet.Currency)
	require.NotZero(t, wallet.ID)
	require.NotZero(t, wallet.CreatedAt)

	return wallet
}

func TestCreateWallet(t *testing.T) {
	user := CreateRandomUser(t)
	CreateRandomWallet(t, user)
}

func TestGetWallet(t *testing.T) {
	user := CreateRandomUser(t)
	wallet1 := CreateRandomWallet(t, user)
	wallet2, err := testQueries.GetWallet(context.Background(), wallet1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, wallet2)

	require.Equal(t, wallet1.ID, wallet2.ID)
	require.Equal(t, wallet1.Name, wallet2.Name)
	require.Equal(t, wallet1.Owner, wallet2.Owner)
	require.Equal(t, wallet1.Currency, wallet2.Currency)
	require.WithinDuration(t, wallet1.CreatedAt, wallet2.CreatedAt, time.Second)
}

func TestDeleteWallet(t *testing.T) {
	user := CreateRandomUser(t)
	wallet1 := CreateRandomWallet(t, user)
	arg := DeleteWalletParams{
		ID:    wallet1.ID,
		Owner: user.Username,
	}
	err := testQueries.DeleteWallet(context.Background(), arg)
	require.NoError(t, err)

	wallet2, err := testQueries.GetWallet(context.Background(), wallet1.ID)
	require.Error(t, err)
	require.True(t, errors.Is(err, sql.ErrNoRows))
	require.Empty(t, wallet2)
}

func TestListWallets(t *testing.T) {
	user := CreateRandomUser(t)
	for i := 0; i < 10; i++ {
		CreateRandomWallet(t, user)
	}

	arg := ListWalletsParams{
		Owner:  user.Username,
		Limit:  5,
		Offset: 5,
	}

	wallets, err := testQueries.ListWallets(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, wallets, 5)
	for _, wallet := range wallets {
		require.NotEmpty(t, wallet)
	}
}
