package db

import (
	"context"
	"testing" 
	"github.com/stretchr/testify/require"
)


func createTestTransfer(t *testing.T) Transfer {
	fromAccount := createTestAccount(t)
	toAccount := createTestAccount(t)
 

	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:  toAccount.ID,
		Amount: 1000,  
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID) 
	require.Equal(t, arg.Amount, transfer.Amount)  
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createTestTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createTestTransfer(t)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID) 
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount) 
}

func TestUpdateTransfer(t *testing.T) {
	transfer1 := createTestTransfer(t)

	arg := UpdateTransferParams{ 
		ID: transfer1.ID,
		FromAccountID: transfer1.FromAccountID,
		ToAccountID: transfer1.ToAccountID,
		Amount: 5000,
	}

	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, arg.ID, transfer2.ID)
	require.Equal(t, arg.FromAccountID, transfer2.FromAccountID) 
	require.Equal(t, arg.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, arg.Amount, transfer2.Amount) 
}

func TestDeleteTransfer(t *testing.T) {
	transfer1 := createTestTransfer(t)

	transfer2, err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount) 

	// Check if the transfer is deleted
	_, err = testQueries.GetTransfer(context.Background(), transfer2.ID)
	require.Error(t, err)
}

func TestListTransfers(t *testing.T) {
	arg := ListTransfersParams{
		Limit:  3,
		Offset: 0,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
	require.Len(t, transfers, 3)

	for _, entry := range transfers {
		require.NotEmpty(t, entry.FromAccountID)
		require.NotEmpty(t, entry.ToAccountID)
		require.NotEmpty(t, entry.Amount) 
	}
}
