package modeltests

import (
	"log"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/MAlifiahM/Vutura_Test_Code/api/models"
	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllTransactions(t *testing.T) {

	err := refreshUserAndTransactionTable()
	if err != nil {
		log.Fatalf("Error refreshing user and transaction table %v\n", err)
	}
	_, _, err = seedUsersAndTransactions()
	if err != nil {
		log.Fatalf("Error seeding user and transaction  table %v\n", err)
	}
	transactions, err := transactionInstance.FindAllTransactions(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the transactions: %v\n", err)
		return
	}
	assert.Equal(t, len(*transactions), 2)
}

func TestSaveTransaction(t *testing.T) {

	err := refreshUserAndTransactionTable()
	if err != nil {
		log.Fatalf("Error user and transaction refreshing table %v\n", err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user %v\n", err)
	}

	newTransaction := models.Transaction{
		ID:       1,
		Product: "Minuman",
		Price: 10000,
		IDUser: user.ID,
	}
	savedTransaction, err := newTransaction.SaveTransaction(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the transaction: %v\n", err)
		return
	}
	assert.Equal(t, newTransaction.ID, savedTransaction.ID)
	assert.Equal(t, newTransaction.Product, savedTransaction.Product)
	assert.Equal(t, newTransaction.Price, savedTransaction.Price)
	assert.Equal(t, newTransaction.IDUser, savedTransaction.IDUser)

}

func TestGetTransactionByID(t *testing.T) {

	err := refreshUserAndTransactionTable()
	if err != nil {
		log.Fatalf("Error refreshing user and transaction table: %v\n", err)
	}
	transaction, err := seedOneUserAndOneTransaction()
	if err != nil {
		log.Fatalf("Error Seeding table")
	}
	foundTransaction, err := transactionInstance.FindTransactionByID(server.DB, transaction.ID)
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, foundTransaction.ID, transaction.ID)
	assert.Equal(t, foundTransaction.Product, transaction.Product)
	assert.Equal(t, foundTransaction.Price, transaction.Price)
}

func TestUpdateATransaction(t *testing.T) {

	err := refreshUserAndTransactionTable()
	if err != nil {
		log.Fatalf("Error refreshing user and post table: %v\n", err)
	}
	transaction, err := seedOneUserAndOneTransaction()
	if err != nil {
		log.Fatalf("Error Seeding table")
	}
	transactionUpdate := models.Transaction{
		ID:       1,
		Product: "Makanan",
		Price: 10000,
		IDUser: transaction.IDUser,
	}
	updatedTransaction, err := transactionUpdate.UpdateATransaction(server.DB)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, updatedTransaction.ID, transactionUpdate.ID)
	assert.Equal(t, updatedTransaction.Product, transactionUpdate.Product)
	assert.Equal(t, updatedTransaction.Price, transactionUpdate.Price)
	assert.Equal(t, updatedTransaction.IDUser, transactionUpdate.IDUser)
}

func TestDeleteATransaction(t *testing.T) {

	err := refreshUserAndTransactionTable()
	if err != nil {
		log.Fatalf("Error refreshing user and transaction table: %v\n", err)
	}
	transaction, err := seedOneUserAndOneTransaction()
	if err != nil {
		log.Fatalf("Error Seeding tables")
	}
	isDeleted, err := transactionInstance.DeleteATransaction(server.DB, transaction.ID, transaction.IDUser)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	//one shows that the record has been deleted or:
	// assert.Equal(t, int(isDeleted), 1)

	//Can be done this way too
	assert.Equal(t, isDeleted, int64(1))
}