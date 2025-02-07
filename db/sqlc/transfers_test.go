package db

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com.victorex27/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func createTestTransfer() Transfer {

	account1 := createTestAccount()
	account2 := createTestAccount()

	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)

	if err != nil {

		log.Fatal("failed to create transfer")
	}

	return transfer

}

func createTestTransfers() []Transfer {

	var transfers []Transfer

	account1 := createTestAccount()
	account2 := createTestAccount()

	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	for i := 0; i < 10; i++ {
		transfer, err := testQueries.CreateTransfer(context.Background(), args)

		if err != nil {

			log.Fatal("failed to create transfer")
		}

		transfers = append(transfers, transfer)
	}

	return transfers

}

func TestCreateTransfer(t *testing.T) {

	account1 := createTestAccount()
	account2 := createTestAccount()

	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, account1.ID, transfer.FromAccountID)
	require.Equal(t, account2.ID, transfer.ToAccountID)
	require.Equal(t, args.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)

}

func TestGetTransfer(t *testing.T) {

	transfer := createTestTransfer()

	retrievedTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, retrievedTransfer)

	require.Equal(t, transfer.ID, retrievedTransfer.ID)
	require.Equal(t, transfer.FromAccountID, retrievedTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, retrievedTransfer.ToAccountID)
	require.Equal(t, transfer.Amount, retrievedTransfer.Amount)

	require.WithinDuration(t, transfer.CreatedAt, retrievedTransfer.CreatedAt, time.Second)

}

func TestUpdateTransfer(t *testing.T) {

	transfer := createTestTransfer()

	arg := UpdateTransferParams{
		ID:     transfer.ID,
		Amount: util.RandomMoney(),
	}

	updatedTransfer, err := testQueries.UpdateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updatedTransfer)

	require.Equal(t, transfer.ID, updatedTransfer.ID)
	require.Equal(t, transfer.FromAccountID, updatedTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, updatedTransfer.ToAccountID)
	require.Equal(t, arg.Amount, updatedTransfer.Amount)

	require.WithinDuration(t, transfer.CreatedAt, updatedTransfer.CreatedAt, time.Second)

}

func TestDeleteTransfer(t *testing.T) {

	transfer := createTestTransfer()

	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)

	retrievedTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, retrievedTransfer)

}

func TestListTransfers(t *testing.T) {

	transfers := createTestTransfers()

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	retrievedTransfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, retrievedTransfers, 5)
	require.Len(t, transfers, 10)

	for _, entry := range retrievedTransfers {

		require.NotEmpty(t, entry)
	}
}


