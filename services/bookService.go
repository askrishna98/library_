package service

import (
	"errors"
	"strings"

	"github.com/askrishna98/library_/models"
)

type BookService struct {
	DB          *models.MockDB
	IdGenerator *IdGenerator
}

func GetInstanceOfBookService(DBInstance *models.MockDB, IdGeneratorInstance *IdGenerator) *BookService {
	return &BookService{
		DB:          DBInstance,
		IdGenerator: IdGeneratorInstance,
	}
}

func (b *BookService) CreateBook(newBook *models.Book) error {

	newBook.Book_id = b.IdGenerator.GenerateBookId()
	b.DB.Books.AddNewItem(newBook)

	return nil
}

// update Booksinfo
func (b *BookService) UpdateBookInfo(bookID int, updateInfo models.BookRequest) (*models.Book, error) {
	book, err := b.BookAvailability(bookID)

	if err != nil {
		return nil, err
	}
	if updateInfo.Title != "" {
		book.Title = updateInfo.Title
	}
	if updateInfo.Author != "" {
		book.Author = updateInfo.Author
	}
	if updateInfo.Category != "" {
		book.Category = updateInfo.Category
	}
	if updateInfo.Count != 0 {
		book.UpdateCount(updateInfo.Count)
	}
	return book, err
}

func (b *BookService) DeleteBook(bookID int) error {

	for i, val := range *b.DB.Books.GetItems() {
		currBook := (*val).(*models.Book)
		if currBook.Book_id == bookID {
			b.DB.Books.DeleteItem(i)
			return nil
		}
	}
	return errors.New("book not found")
}

// to know book is available or present  in library
func (b *BookService) BookAvailability(bookID int) (*models.Book, error) {

	for _, val := range *b.DB.Books.GetItems() {
		currBook := (*val).(*models.Book)
		if currBook.Book_id == bookID {
			return currBook, nil
		}
	}
	return nil, errors.New("book not found")
}

// to makesure count > 0
func (b *BookService) BookCount(book *models.Book) error {
	if book.Count > 0 {
		return nil
	} else {
		return errors.New("the book is currently unavailable")
	}
}

// to get list of books by Category,authorname,prefix
// new filter func
func (b *BookService) Filter(author, category, prefix string) []models.BookResponse {

	filtered_books := []models.BookResponse{}

	for _, val := range *b.DB.Books.GetItems() {
		book := (*val).(*models.Book)
		if authorMatches(book.Author, author) && categoryMatches(book.Category, category) && prefixMatches(book.Title, prefix) {
			currBook := models.BookResponse{
				Book_id:  book.Book_id,
				Title:    book.Title,
				Author:   book.Author,
				Category: book.Category,
			}
			filtered_books = append(filtered_books, currBook)
		}
	}
	return filtered_books
}

func authorMatches(bookAuthor, author string) bool {
	return author == "" || strings.EqualFold(author, bookAuthor)
}

func categoryMatches(bookCategory, category string) bool {
	return category == "" || strings.EqualFold(category, bookCategory)
}

func prefixMatches(bookTitle, prefix string) bool {
	return prefix == "" || (len(bookTitle) >= len(prefix) && strings.EqualFold(bookTitle[:len(prefix)], prefix))
}
