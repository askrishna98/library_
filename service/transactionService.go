package service

import (
	"time"

	"github.com/askrishna98/library_/models"
)

type TransactionService struct {
	DB                    *models.MockDB
	MemberServiceInstance *MemberService // uses some methods from Member Service as well as Book service
	BookServiceInstance   *BookService
}

func GetInstanceOfTransactionService(DBInstance *models.MockDB,
	MemberServicePointer *MemberService,
	BookServicePointer *BookService) *TransactionService {
	return &TransactionService{
		DB:                    DBInstance,
		MemberServiceInstance: MemberServicePointer,
		BookServiceInstance:   BookServicePointer,
	}
}

// Creating a new Book transaction
func (t TransactionService) BorrowBook(memberID string, bookID int) error {

	member, err := t.MemberServiceInstance.GetMemberById(memberID)

	if err != nil {
		return err
	}

	book, err := t.BookServiceInstance.BookAvailability(bookID)

	if err != nil {
		return err
	}

	var NewTransaction = models.Transaction{
		Borrow_id:   len(t.DB.BookTransactions) + 1,
		Member:      member,
		Book:        book,
		Borrow_date: time.Now().Format("02-01-2006"),
		Due_date:    time.Now().AddDate(0, 0, 10).Format("02-01-2006"),
	}

	t.DB.BookTransactions = append(t.DB.BookTransactions, NewTransaction)

	return nil
}


// TO return the BOOk
func (t TransactionService) ReturnBook(memberID string, bookID int) error {
	member, err := t.MemberServiceInstance.GetMemberById(memberID)

	if err != nil {
		return err
	}

	book, err := t.BookServiceInstance.BookAvailability(bookID)

	if err != nil {
		return err
	}

	for _, transaction := range t.DB.BookTransactions {
		if transaction.Return_date == "" && transaction.Member == member && transaction.Book == book {
			transaction.Return_date = time.Now().Format("02-01-2006")
			return nil
		}
	}
	return nil
}
