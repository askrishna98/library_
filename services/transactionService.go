package service

import (
	"sync"
	"time"

	"github.com/askrishna98/library_/models"
)

type TransactionService struct {
	DB                    *models.MockDB
	IdGenerator           *IdGenerator
	MemberServiceInstance *MemberService // uses some methods from Member Service as well as Book service
	BookServiceInstance   *BookService
	mutex                 *sync.Mutex
}

func GetInstanceOfTransactionService(DBInstance *models.MockDB,
	IdGeneratorInstance *IdGenerator,
	MemberServicePointer *MemberService,
	BookServicePointer *BookService) *TransactionService {
	return &TransactionService{
		DB:                    DBInstance,
		MemberServiceInstance: MemberServicePointer,
		BookServiceInstance:   BookServicePointer,
		IdGenerator:           IdGeneratorInstance,
		mutex:                 &sync.Mutex{},
	}
}

// Creating a new Book transaction
func (t *TransactionService) BorrowBook(memberID string, bookID int) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	member, err := t.MemberServiceInstance.GetMemberById(memberID)

	if err != nil {
		return err
	}

	book, err := t.BookServiceInstance.BookAvailability(bookID)

	if err != nil {
		return err
	}

	if err := t.BookServiceInstance.BookCount(book); err != nil {
		return err
	}

	var NewTransaction = models.Transaction{
		Borrow_id:   t.IdGenerator.GenerateTransactionId(),
		Member:      member,
		Book:        book,
		Borrow_date: time.Now().Format("02-01-2006"),
		Due_date:    time.Now().AddDate(0, 0, 10).Format("02-01-2006"),
	}
	// decrease book count
	book.Count--

	t.DB.BookTransactions = append(t.DB.BookTransactions, &NewTransaction)

	return nil
}

// TO return the BOOk
func (t *TransactionService) ReturnBook(memberID string, bookID int) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	member, err := t.MemberServiceInstance.GetMemberById(memberID)

	if err != nil {
		return err
	}

	book, err := t.BookServiceInstance.BookAvailability(bookID)

	if err != nil {
		return err
	}

	for _, transaction := range t.DB.BookTransactions {
		if transaction.Return_date == "" && transaction.Member.Member_id == member.Member_id && transaction.Book.Book_id == book.Book_id {
			transaction.Return_date = time.Now().Format("02-01-2006")
			book.Count++
			return nil
		}
	}

	return nil
}
