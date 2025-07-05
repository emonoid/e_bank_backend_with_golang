package db

import (
	"context"
	"testing"
	// "time"  

	// "github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)


func createTestAccount(t *testing.T) Account {
	userArg := CreateUserParams{ 
		 Username: "emonoid",
		 HashedPassword: "secret",
		 FullName: "Al Sahriar",
		 Email: "emon@gmail.com",
	}
   user, err := testQueries.CreateUser(context.Background(), userArg)
   require.NoError(t, err)

	// createdAt := pgtype.Timestamptz{}
	// _ = createdAt.Scan(time.Now())

	arg := CreateAccountParams{ 
		Owner: user.Username,
		Balance:   500,
		Currency:  "BDT", 
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	// require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createTestAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createTestAccount(t)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createTestAccount(t)

	arg := UpdateAccountParams{ 
		ID: account1.ID,
		Owner: "Updated Owner",
		Balance:   500,
		Currency:  "USD",
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, arg.ID, account2.ID)
	require.Equal(t, arg.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, arg.Currency, account2.Currency)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createTestAccount(t)

	account2, err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)

	// Check if the account is deleted
	_, err = testQueries.GetAccount(context.Background(), account1.ID) 
	require.Error(t, err)
}

func TestListAccounts(t *testing.T) {
	arg := ListAccountsParams{
		Limit:  3,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	require.Len(t, accounts, 3)

	for _, account := range accounts {
		require.NotEmpty(t, account.Owner)
		require.NotEmpty(t, account.Balance)
		require.NotEmpty(t, account.Currency)
	}
}
