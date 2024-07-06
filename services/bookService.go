package service

import (
	"errors"
	"strings"
	"sync"

	"github.com/askrishna98/library_/models"
)

type BookService struct {
	DB          *models.MockDB
	IdGenerator *IdGenerator
	mutex       *sync.Mutex
}

func GetInstanceOfBookService(DBInstance *models.MockDB, IdGeneratorInstance *IdGenerator) *BookService {
	return &BookService{
		DB:          DBInstance,
		IdGenerator: IdGeneratorInstance,
		mutex:       &sync.Mutex{},
	}
}

func (b *BookService) CreateBook(newBook *models.Book) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	newBook.Book_id = b.IdGenerator.GenerateBookId()
	b.DB.Books = append(b.DB.Books, newBook)
	return nil
}

func (b *BookService) DeleteBook(bookID int) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

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
	b.mutex.Lock()
	defer b.mutex.Unlock()

	for _, book := range b.DB.Books {
		if book.Book_id == bookID {
			return book, nil
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

// new filter func
func (b *BookService) Filter(author, category, prefix string) []models.Book {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	filtered_books := []models.Book{}

	for _, book := range b.DB.Books {
		if authorMatches(book.Author, author) && categoryMatches(book.Category, category) && prefixMatches(book.Title, prefix) {
			filtered_books = append(filtered_books, *book)
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
