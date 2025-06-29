package db

import (
	"context"
	"testing" 
 
	"github.com/stretchr/testify/require"
)


func createTestEntry(t *testing.T) Entry {
	account := createTestAccount(t) 

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount: 1000, 
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)  
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createTestEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createTestEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount) 
}

func TestUpdateEntry(t *testing.T) {
	entry1 := createTestEntry(t)

	arg := UpdateEntryParams{ 
		ID: entry1.ID,
		AccountID: 1,
		Amount: 102, 
	}

	entry2, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, arg.ID, entry2.ID)
	require.Equal(t, arg.AccountID, entry2.AccountID)
	require.Equal(t, arg.Amount, entry2.Amount) 
}

func TestDeleteEntry(t *testing.T) {
	entry1 := createTestEntry(t)

	entry2, err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount) 

	// Check if the entry is deleted
	_, err = testQueries.GetEntry(context.Background(), entry2.ID)
	require.Error(t, err)
}

func TestListEntries(t *testing.T) {
	arg := ListEntryParams{
		Limit:  3,
		Offset: 0,
	}

	entries, err := testQueries.ListEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)
	require.Len(t, entries, 3)

	for _, entry := range entries {
		require.NotEmpty(t, entry.AccountID)
		require.NotEmpty(t, entry.Amount) 
	}
}
