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

func (b BookService) CreateBook(newBook *models.Book) error {
	newBook.Book_id = b.IdGenerator.GenerateBookId()
	b.DB.Books = append(b.DB.Books, newBook)
	return nil
}

func (b BookService) DeleteBook(bookID int) error {
	for i, book := range b.DB.Books {
		if book.Book_id == bookID {
			b.DB.Books = append(b.DB.Books[:i], b.DB.Books[:i+1]...)
			return nil
		}
	}
	return errors.New("book not found")
}

// to know book is available or present  in library
func (b *BookService) BookAvailability(bookID int) (*models.Book, error) {
	for _, book := range b.DB.Books {
		if book.Book_id == bookID {
			if book.Count > 0 {
				return book, nil
			} else {
				return nil, errors.New("the book is currently unavailable")
			}
		}
	}
	return nil, errors.New("book not found")
}

// to get list of books by author name
func (b *BookService) GetBooksByAuthor(author string) []models.Book {

	related_books := []models.Book{}

	for _, book := range b.DB.Books {
		if strings.EqualFold(book.Author, author) {
			related_books = append(related_books, *book)
		}
	}

	return related_books
}

// to get list of books by category
func (b *BookService) GetBooksByCategory(category string) []models.Book {

	related_books := []models.Book{}

	for _, book := range b.DB.Books {
		if strings.EqualFold(book.Category, category) {
			related_books = append(related_books, *book)
		}
	}

	return related_books
}

// to get list of books by prefix

func (b *BookService) GetBooksByPrefix(prefix string) []models.Book {

	related_books := []models.Book{}

	for _, book := range b.DB.Books {
		if len(book.Title) >= len(prefix) && strings.EqualFold(book.Title[:len(prefix)], prefix) {
			related_books = append(related_books, *book)
		}
	}
	return related_books
}
