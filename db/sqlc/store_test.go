package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTrxn(t *testing.T){
	store := NewStore(testConn)

	account1 := createTestAccount(t)
	account2 := createTestAccount(t)

	// run concurrent transfers
	n := 5
	amount := int64(10)

	// channels
	errs := make(chan error)
	results := make(chan TransferTrxnResult)

	for i:=0; i<n; i++{
		go func ()  {
			result, err := store.TransferTrnx(context.Background(), TransferTrxnParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// check results from channels
	for i := 0; i < n; i++ {
		err:= <-errs
		require.NoError(t, err)

		result := <-results
		
		// check the transfer result
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, transfer.FromAccountID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)
		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, transfer.ToAccountID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)
		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// check accounts
		// TODO: check account balances
		// ...
	}

}