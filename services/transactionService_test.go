package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/askrishna98/library_/models"
	"github.com/stretchr/testify/assert"
)

func TestTransactionService(t *testing.T) {
	assert := assert.New(t)
	MockDB := models.GetMockDBInstance()
	IDgen := InitalizeIDGenerator()
	BookService := GetInstanceOfBookService(MockDB, IDgen)
	MemberService := GetInstanceOfMemberService(MockDB, IDgen)
	TransactionService := GetInstanceOfTransactionService(MockDB, IDgen, MemberService, BookService)

	//adding some test values in mockDB
	BookService.CreateBook(&models.Book{Title: "book1", Author: "author1", Category: "cat1", Count: 1})
	MemberService.CreateMember(&models.Member{Name: "member1", Phone: "123456789", Email: "member@example.com"})

	t.Run("BorrowBook_Success", func(t *testing.T) {
		memberID, bookID := "A001", 1
		Transaction, err := TransactionService.BorrowBook(memberID, bookID)
		assert.NoError(err)
		assert.Equal(Transaction.Book.Book_id, bookID, "Book_id dont match")
		assert.Equal(Transaction.Member.Member_id, memberID, "Member_ID dont match")
		assert.Equal(len(MockDB.BookTransactions), 1, fmt.Sprintf("expected length of booktransacction size as 1 but got %d", len(MockDB.BookTransactions)))
		// Checking decrease in Book Count
		assert.Equal(Transaction.Book.Count, 0)

	})
	t.Run("BorrowBook_fail", func(t *testing.T) {
		// Invalid Book ID
		_, err := TransactionService.BorrowBook("A002", 2)
		assert.NotNil(err)
		// Invalid Member ID
		_, err = TransactionService.BorrowBook("A001", 1)
		assert.NotNil(err)
	})

	// Borrowing Book which have count 0
	t.Run("BorrowBok_CountZero", func(t *testing.T) {
		// adding a new book To mockDB with zero count
		BookService.CreateBook(&models.Book{Title: "book2", Author: "author2", Category: "cat2", Count: 0})

		_, err := TransactionService.BorrowBook("A001", 2)
		assert.NotNil(err, "expected Error 'book not available'")
	})

	// Tests for returning Book
	t.Run("ReturnBook_Success", func(t *testing.T) {
		memberID, BookID, date := "A001", 1, time.Now().Format("02-01-2006")
		Transaction, Penalty, err := TransactionService.ReturnBook(memberID, BookID)
		assert.NotNil(Transaction)
		assert.Nil(err)
		assert.Equal(Penalty, 0)
		assert.Equal(Transaction.Book.Book_id, BookID, "BOokID do not match")
		assert.Equal(Transaction.Member.Member_id, memberID, "MemberID do not match")
		assert.NotEmpty(Transaction.Return_date, "return_date field should not be empty")
		assert.Equal(Transaction.Return_date, date)
		// check Increase in BookCount
		assert.Equal(Transaction.Book.Count, 1)
	})

	t.Run("ReturnBook_Failed", func(t *testing.T) {
		// Invalid BookID
		_, _, err := TransactionService.ReturnBook("A001", 3)
		assert.NotNil(err)
		// Invalid MemberID
		_, _, err = TransactionService.ReturnBook("A002", 1)
		assert.NotNil(err)
		//No matching entry - like (adding a new book ID=3 but "A001" never borrowed that Book)
		BookService.CreateBook(&models.Book{Title: "book3", Author: "author3", Category: "cat3", Count: 5})
		_, _, err = TransactionService.ReturnBook("A001", 3)
		assert.NotNil(err)
	})

	t.Run("TestCalPenalty", func(t *testing.T) {
		tests := []struct {
			Borrowdate      string
			Returndate      string
			ExpectedPenalty int
		}{
			{"05-07-2024", "10-07-2024", 0},
			{"01-07-2024", "11-07-2024", 0},
			{"01-07-2024", "12-07-2024", 50},
			{"15-06-2024", "10-07-2024", 100},
			{"10-05-2024", "10-07-2024", 300},
		}

		for _, test := range tests {
			penalty := Calpenalty(test.Borrowdate, test.Returndate)
			assert.Equal(test.ExpectedPenalty, penalty)
		}
	})
	t.Run("BorrowBook_With_Penalty", func(t *testing.T) {
		// adding some testValues
		MockDB.Members = []*models.Member{
			{Member_id: "A001", Name: "member1", Phone: "123456789", Email: "member@example.com"},
		}
		MockDB.Books = []*models.Book{
			{Book_id: 1, Title: "book1", Author: "author1", Category: "cat1", Count: 1},
		}
		MockDB.BookTransactions = []*models.Transaction{}

		MockDB.BookTransactions = append(MockDB.BookTransactions, &models.Transaction{
			Borrow_id:   1,
			Member:      MockDB.Members[0],
			Book:        MockDB.Books[0],
			Borrow_date: "01-07-2024",
			Due_date:    "10-07-2024",
		})

		Transaction, penalty, err := TransactionService.ReturnBook("A001", 1)
		assert.Nil(err)
		assert.NotNil(Transaction)
		assert.Equal(50, penalty)
	})
}
