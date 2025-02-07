package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createTestAccount()
	account2 := createTestAccount()

	n := 3
	amount := int64(10)

	errs := make(chan error, n)
	results := make(chan TransferTxResult, n)
	amounts := make(chan int64, n)

	for i := 0; i < n; i++ {
		// use this name to debug deadlock
		txName := fmt.Sprintf("tx %d", i+1)
		go func() {

			
			txKey := transactionKey(txName) 

			ctx := context.WithValue(context.Background(), txKey, txName)

			result, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
				TxKey: txKey,
			})

			errs <- err
			results <- result
			amounts <- amount

			fmt.Println("ending go routine")
		}()
	}

	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		amount := <-amounts
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// check fromEntry
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		// check toEntry
		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// check account for sender
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.NotZero(t, fromAccount.ID)
		require.NotZero(t, fromAccount.CreatedAt)

		// check account for recipient
		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.NotZero(t, toAccount.ID)
		require.NotZero(t, toAccount.CreatedAt)

		// check balance
		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance

		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0) // the diff should be divisible by the amount

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)

		existed[k] = true

		fmt.Println("After tx:", fromAccount.Balance, toAccount.Balance)

		

	}

	// check updated balance
	updatedAccount1, err := store.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := store.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after: acct 1 ", updatedAccount1.Balance, account1.Balance-int64(n)*amount)
	fmt.Println(">> after: acct 2 ", updatedAccount2.Balance, account2.Balance+int64(n)*amount)
	fmt.Println(">> amount ", amount)
	fmt.Println(">> account 1 balance ", account1.Balance)
	fmt.Println(">> account 2 balance ", account2.Balance)

	require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
	// require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)

}

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)

	account1 := createTestAccount()
	account2 := createTestAccount()

	// this will only work for even numbers as we are doing the same
	// number of transaction for each user
	n := 6

	errs := make(chan error, n)

	amount := int64(10)

	for i := 0; i < n; i++ {

		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}

		// use this name to debug deadlock
		txName := fmt.Sprintf("tx %d", i+1)
		txKey := transactionKey(txName) 
		go func() {

			ctx := context.WithValue(context.Background(), txKey, txName)

			_, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
				TxKey: txKey,
			})

			errs <- err

			

			fmt.Println("ending go routine")
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

	}

	//check the final updated balance
	updatedAccount1, err := store.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := store.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)

	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)

}
