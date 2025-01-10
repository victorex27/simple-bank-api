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

func createTestAccount() Account{

	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount( context.Background(), arg)

	if err != nil {
		log.Fatal("did not create an account")
	}

	return account
}


func TestCreateAccount(t *testing.T){
	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}


	account, err := testQueries.CreateAccount( context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)


	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)


	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

}

func TestGetAccount(t *testing.T){

	account := createTestAccount()
	


	retrievedAccount,err := testQueries.GetAccount(context.Background(),account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, retrievedAccount)
	
	require.Equal(t, account.ID, retrievedAccount.ID)
	require.Equal(t, account.Owner, retrievedAccount.Owner)
	require.Equal(t, account.Balance, retrievedAccount.Balance)
	require.Equal(t, account.Currency, retrievedAccount.Currency)

	require.WithinDuration(t, account.CreatedAt, retrievedAccount.CreatedAt, time.Second)

}


func TestUpdateAccount(t *testing.T){

	

	account := createTestAccount()



	arg := UpdateAccountParams{
		ID: account.ID,
		Balance: util.RandomMoney(),
	}
	


	updatedAccount,err := testQueries.UpdateAccount(context.Background(), arg)



	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)
	
	require.Equal(t, account.ID, updatedAccount.ID)
	require.Equal(t, account.Owner, updatedAccount.Owner)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, account.Currency, updatedAccount.Currency)

	require.WithinDuration(t, account.CreatedAt, updatedAccount.CreatedAt, time.Second)

}


func TestDeleteAccount(t *testing.T){

	

	account := createTestAccount()
	

	err := testQueries.DeleteAccount(context.Background(), account.ID)

	require.NoError(t, err)

	retrievedAccount,err := testQueries.GetAccount(context.Background(), account.ID)
	

	require.Error(t,err)
	require.EqualError(t,err, sql.ErrNoRows.Error())
	require.Empty(t, retrievedAccount)
	

}


func TestListAccounts(t *testing.T){

	for i := 0; i < 10; i++ {
		createTestAccount()
	}
	// account := createTestAccount()
	
	arg := ListAccountsParams{
		Limit: 5,
		Offset: 5,
	}

	retrievedAccount,err := testQueries.ListAccounts(context.Background(),arg)

	require.NoError(t, err)
	require.Len(t, retrievedAccount, 5)
	
	for _, account := range retrievedAccount{

		require.NotEmpty(t, account)
	}
}