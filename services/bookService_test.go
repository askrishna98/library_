package service

import (
	"testing"

	"github.com/askrishna98/library_/models"
	"github.com/stretchr/testify/assert"
)

func TestBookService(t *testing.T) {
	mockDB := models.GetMockDBInstance()
	IdGenerator := InitalizeIDGenerator()
	BookService := GetInstanceOfBookService(mockDB, IdGenerator)
	assert := assert.New(t)

	t.Run("CreateBook", func(t *testing.T) {
		testbook := &models.Book{
			Title:    "book1",
			Author:   "author1",
			Category: "cat1",
			Count:    1,
		}
		err := BookService.CreateBook(testbook)
		assert.NoError(err)
		assert.Equal(1, testbook.Book_id, "bookId do not match")
		assert.Equal(1, len(mockDB.Books), "length of bookslice in DB missmatch")
	})

	t.Run("BookAvailability", func(t *testing.T) {

		book, err := BookService.BookAvailability(1)
		assert.NoError(err)
		assert.Equal(book.Book_id, 1, "BookiD missmatch")
	})

	t.Run("BookCount", func(t *testing.T) {
		book, err := BookService.BookAvailability(1)
		assert.NoError(err)
		err = BookService.BookCount(book)
		assert.NoError(err)
	})

	t.Run("FilterBooks", func(t *testing.T) {
		mockDB.Books = []*models.Book{
			{Title: "The Hobbit", Author: "J.R.R. Tolkien", Category: "Fantasy", Count: 8},
			{Title: "Crime and Punishment", Author: "Fyodor Dostoevsky", Category: "Psychological", Count: 3},
			{Title: "Brave New World", Author: "Aldous Huxley", Category: "Dystopian", Count: 4},
			{Title: "The Odyssey", Author: "Homer", Category: "Epic", Count: 6},
			{Title: "Jane Eyre", Author: "Charlotte Bronte", Category: "Fantasy", Count: 3},
		}
		tests := []struct {
			name          string
			author        string
			category      string
			prefix        string
			expectedCount int
		}{
			{"filter by title author", "J.R.R. Tolkien", "", "", 1},
			{"filter by title category", "", "Fantasy", "", 2},
			{"filter by title prefix", "", "", "Brave", 1},
			{"filter by title author and category", "J.R.R. Tolkien", "Fantasy", "", 1},
			{"filter by All", "Fyodor Dostoevsky", "Psychological", "Crime", 1},
			{"filter by nothing - gets all books", "", "", "", 5},
		}

		for _, tc := range tests {
			filteredBooks := BookService.Filter(tc.author, tc.category, tc.prefix)
			assert.Equal(tc.expectedCount, len(filteredBooks), tc.name+" - Count do not match")
		}
	})

	t.Run("DeleteBook", func(t *testing.T) {
		mockDB.Books = []*models.Book{
			{Book_id: 1, Title: "Book1", Author: "Author1", Category: "Category1", Count: 8},
		}
		err := BookService.DeleteBook(1)
		assert.NoError(err)
		assert.Equal(len(mockDB.Books), 0)
	})
}
