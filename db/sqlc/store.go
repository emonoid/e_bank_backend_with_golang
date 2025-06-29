package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute SQL queries and transactions.
// It wraps the generated Queries type and provides a database connection.
type Store struct {
	*Queries         // embedded Queries pointer as like inheritance
	db       *sql.DB // same
}

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db), // initialize Queries with the provided db connection
		db:      db,      // store the db connection
	}
}

// execTrxn executes a function within a transaction context.
// It begins a transaction, executes the provided function with a Queries instance,
// and commits the transaction if successful. If an error occurs, it rolls back the transaction.
func (store *Store) execTrxn(ctx context.Context, fn func(*Queries) error) error {
	trxn, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(trxn)
	err = fn(q)

	if err != nil {
		if rollbackError := trxn.Rollback(); rollbackError != nil {
			return fmt.Errorf("transaction rollback error: %v, original error: %v", rollbackError, err)
		}

		return err
	}

	return trxn.Commit()
}

// TransferTrxnParams contains the parameters for the TransferTrnx function.
type TransferTrxnParams struct {
	FromAccountID int64 `json: from_accoount_id`
	ToAccountID   int64 `json: to_account_id`
	Amount        int64 `json: amount`
}

// TransferTrxnResult contains the result of the TransferTrnx function.
type TransferTrxnResult struct {
	FromAccount Account  `json: from_account`
	ToAccount   Account  `json: to_account`
	Transfer    Transfer `json: transfer`
	FromEntry   Entry    `json: from_entry`
	ToEntry     Entry    `json: to_entry`
}

// TransferTrxn performs a money transfer between two accounts.
// It uses a transaction to ensure atomicity, meaning either both operations succeed or none do.
// It creates a transfer record, updates the account balances, and creates entries for both accounts.
func (store *Store) TransferTrxn(ctx context.Context, arg TransferTrxnParams) (TransferTrxnResult, error) {
	var result TransferTrxnResult

	err := store.execTrxn(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID, ToAccountID: arg.ToAccountID, Amount: arg.Amount})

		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// update account balances

		if arg.FromAccountID < arg.ToAccountID { // to prevent deadlock, must insert same id at first operation of all trxns at concurrent, so we start from lowest id
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountID, arg.ToAccountID, -arg.Amount, arg.Amount)
			if err != nil {
				return err
			}
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountID, arg.FromAccountID, arg.Amount, -arg.Amount)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return result, err
}

func addMoney(ctx context.Context, q *Queries, accountID1 int64, accountID2 int64, amount1 int64, amount2 int64) (account1 Account, account2 Account, err error) {
	account1, err = q.AddBalance(ctx, AddBalanceParams{
		ID:     accountID1,
		Amount: amount1})

	if err != nil {
		return
	}

	account2, err = q.AddBalance(ctx, AddBalanceParams{
		ID:     accountID2,
		Amount: amount2})

	return
}
