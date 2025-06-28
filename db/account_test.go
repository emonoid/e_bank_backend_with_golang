package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createTestAccount(t *testing.T, id int64) Account {

	createdAt := pgtype.Timestamptz{}
	_ = createdAt.Scan(time.Now())

	arg := CreateAccountParams{
		ID: id,
		OwnerName: "Al Sahriar",
		Balance:   12300000,
		Currency:  "BDT",
		CreatedAt: createdAt,
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.OwnerName, account.OwnerName)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	// require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createTestAccount(t, 15)
}

func TestGetAccount(t *testing.T) {
	account1 := createTestAccount(t, 16)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.OwnerName, account2.OwnerName)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createTestAccount(t, 17)

	arg := UpdateAccountParams{ 
		ID: account1.ID,
		OwnerName: "Updated Owner",
		Balance:   15000000,
		Currency:  "USD",
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, arg.ID, account2.ID)
	require.Equal(t, arg.OwnerName, account2.OwnerName)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, arg.Currency, account2.Currency)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createTestAccount(t, 18)

	account2, err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.OwnerName, account2.OwnerName)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)

	// Check if the account is deleted
	_, err = testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
}
