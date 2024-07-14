package service

import (
	"testing"

	"github.com/askrishna98/library_/models"
	service "github.com/askrishna98/library_/services"
	"github.com/stretchr/testify/assert"
)

func TestTransactionService(t *testing.T) {
	assert.New(t)
	MockDB := models.GetMockDBInstance()
	IDgen := InitalizeIDGenerator()
	BookService := GetInstanceOfBookService(MockDB,IDgen)
	MemberService := GetInstanceOfMemberService(MockDB,IDgen)
	TransactionService := GetInstanceOfTransactionService(MockDB,IDgen,MemberService,BookService)

	t.Run("BorrowBook",func(t *testing.T) {
		
	})

}
