package service

import (
	"errors"
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

// Performs a List of Book transaction
func (t *TransactionService) BorrowBooks(memberID string, bookIDs []int) (*models.BookTransactionResponse, error) {
	res := &models.BookTransactionResponse{}
	member, err := t.MemberServiceInstance.GetMemberById(memberID)

	if err != nil {
		return nil, err
	}

	res.Member = struct {
		MemberID string
		Name     string
	}{member.Member_id, member.Name}

	for _, id := range bookIDs {
		trans, err := t.BorrowBook(memberID, id)
		if err != nil {
			res.Books = append(res.Books, struct {
				BookID       int
				Title        string
				Borrowdate   string
				DueDate      string
				ErrorMessage string
			}{0, "nil", "nil", "nil", err.Error()})
		} else {
			res.Books = append(res.Books, struct {
				BookID       int
				Title        string
				Borrowdate   string
				DueDate      string
				ErrorMessage string
			}{
				BookID:       trans.Book.Book_id,
				Title:        trans.Book.Title,
				Borrowdate:   trans.Borrow_date,
				DueDate:      trans.Due_date,
				ErrorMessage: "nil"})
		}
	}
	return res, nil
}

// a single book
func (t *TransactionService) BorrowBook(memberID string, bookID int) (*models.Transaction, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	member, err := t.MemberServiceInstance.GetMemberById(memberID)

	if err != nil {
		return nil, err
	}

	book, err := t.BookServiceInstance.BookAvailability(bookID)

	if err != nil {
		return nil, err
	}

	if err := t.BookServiceInstance.BookCount(book); err != nil {
		return nil, err
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

	return &NewTransaction, nil
}

func (t *TransactionService) ReturnBooks(memberID string, bookIDs []int) (*models.BookReturnTransactionResponse, error) {
	res := &models.BookReturnTransactionResponse{}
	member, err := t.MemberServiceInstance.GetMemberById(memberID)

	if err != nil {
		return nil, err
	}

	res.Member = struct {
		MemberID string
		Name     string
	}{member.Member_id, member.Name}

	for _, bookID := range bookIDs {
		trans, penalty, err := t.ReturnBook(memberID, bookID)
		if err != nil {
			res.Books = append(res.Books, struct {
				BookID        int
				Title         string
				Borrowdate    string
				DueDate       string
				Returned_Date string
				Penalty       int
				ErrorMessage  string
			}{0, "nil", "nil", "nil", "nil", 0, err.Error()})
		} else {
			res.Books = append(res.Books, struct {
				BookID        int
				Title         string
				Borrowdate    string
				DueDate       string
				Returned_Date string
				Penalty       int
				ErrorMessage  string
			}{
				BookID:        trans.Book.Book_id,
				Title:         trans.Book.Title,
				Borrowdate:    trans.Borrow_date,
				DueDate:       trans.Due_date,
				Returned_Date: trans.Return_date,
				Penalty:       penalty,
				ErrorMessage:  "nil",
			})
		}
	}
	return res, nil

}

// TO return a single BOOk
func (t *TransactionService) ReturnBook(memberID string, bookID int) (*models.Transaction, int, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	member, err := t.MemberServiceInstance.GetMemberById(memberID)

	if err != nil {
		return nil, 0, err
	}

	book, err := t.BookServiceInstance.BookAvailability(bookID)

	if err != nil {
		return nil, 0, err
	}

	for _, transaction := range t.DB.BookTransactions {
		if transaction.Return_date == "" && transaction.Member.Member_id == member.Member_id && transaction.Book.Book_id == book.Book_id {
			transaction.Return_date = time.Now().Format("02-01-2006")
			transaction.Book.Count++
			return transaction, Calpenalty(transaction.Borrow_date, transaction.Return_date), err
		}
	}

	return nil, 0, errors.New("NO matched entries")
}

func Calpenalty(Borrow_date, Return_date string) int {
	const penalty int = 50
	const days int = 10
	Bdate, _ := time.Parse("02-01-2006", Borrow_date)
	Rdate, _ := time.Parse("02-01-2006", Return_date)
	difference := int(Rdate.Sub(Bdate).Hours() / 24)
	if difference > 10 {
		return int(difference/days) * penalty
	}
	return 0
}

// Get list of Books borrowed by a Member

func (t *TransactionService) GetBooksBorrowedByMember(memberID string) ([]models.Book, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	Books := []models.Book{}

	member, err := t.MemberServiceInstance.GetMemberById(memberID)
	if err != nil {
		return Books, err
	}

	for _, transaction := range t.DB.BookTransactions {
		if transaction.Member == member && transaction.Return_date == "" {
			Books = append(Books, *transaction.Book)
		}
	}
	return Books, nil
}

// Upcoming return Books
func (t *TransactionService) UpcomingReturns(date string) []models.UpcomingReturnsResponse {
	returns := []models.UpcomingReturnsResponse{}

	if date == "" {
		date = "31-12-9999"
	}
	layout := "02-01-2006"
	dateframe, _ := time.Parse(layout, date)
	for _, transaction := range t.DB.BookTransactions {

		dueDate, _ := time.Parse(layout, transaction.Due_date)

		if !dueDate.After(dateframe) {
			entry := models.UpcomingReturnsResponse{
				MemberID:   transaction.Member.Member_id,
				MemberName: transaction.Member.Name,
				BookID:     transaction.Book.Book_id,
				Title:      transaction.Book.Title,
				BorrowDate: transaction.Borrow_date,
				DueDate:    transaction.Due_date,
			}
			returns = append(returns, entry)
		}
	}
	return returns
}
