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

func createTestEntry() (Entry, Account) {

	account := createTestAccount()

	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)

	if err != nil {

		log.Fatal("failed to create entry")
	}

	return entry, account

}

func createTestEntries() []Entry {

	var entries []Entry
	account := createTestAccount()

	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	for i := 0; i < 10; i++ {
		entry, err := testQueries.CreateEntry(context.Background(), args)

		if err != nil {

			log.Fatal("failed to create entry")
		}

		entries = append(entries, entry)
	}

	return entries

}

func TestCreateEntry(t *testing.T) {

	account := createTestAccount()

	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, account.ID, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)

	require.NotZero(t, entry.AccountID)

}

func TestGetEntry(t *testing.T) {

	entry, account := createTestEntry()

	retrievedEntry, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, retrievedEntry)

	require.Equal(t, entry.ID, retrievedEntry.ID)
	require.Equal(t, entry.AccountID, retrievedEntry.AccountID)
	require.Equal(t, account.ID, retrievedEntry.AccountID)
	require.Equal(t, entry.Amount, retrievedEntry.Amount)

	require.WithinDuration(t, entry.CreatedAt, retrievedEntry.CreatedAt, time.Second)

}

func TestUpdateEntry(t *testing.T) {

	entry, account := createTestEntry()

	arg := UpdateEntryParams{
		ID:     entry.ID,
		Amount: util.RandomMoney(),
	}

	updatedEntry, err := testQueries.UpdateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updatedEntry)

	require.Equal(t, entry.ID, updatedEntry.ID)
	require.Equal(t, entry.AccountID, updatedEntry.AccountID)
	require.Equal(t, arg.Amount, updatedEntry.Amount)
	require.Equal(t, account.ID, updatedEntry.AccountID)

	require.WithinDuration(t, entry.CreatedAt, updatedEntry.CreatedAt, time.Second)

}

func TestDeleteEntry(t *testing.T) {

	entry, _ := createTestEntry()

	err := testQueries.DeleteEntry(context.Background(), entry.ID)

	require.NoError(t, err)

	retrievedAccount, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, retrievedAccount)

}

func TestListEntries(t *testing.T) {

	entries := createTestEntries()

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	retrievedEntries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, retrievedEntries, 5)
	require.Len(t, entries, 10)

	for _, entry := range retrievedEntries {

		require.NotEmpty(t, entry)
	}
}
