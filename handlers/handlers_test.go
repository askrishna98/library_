package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/askrishna98/library_/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandlers(t *testing.T) {
	assert := assert.New(t)

	router := gin.Default()
	Handlers(router)

	// Test for Member routes
	t.Run("Test_Create_MemberAPI", func(t *testing.T) {

		testMember := &models.Member{
			Name:  "test",
			Phone: "123456789",
			Email: "test@example",
		}

		memberJson, _ := json.Marshal(testMember)

		req, _ := http.NewRequest("POST", "/api/members", strings.NewReader(string(memberJson)))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(http.StatusOK, w.Code)
		var respose models.Member
		err := json.Unmarshal(w.Body.Bytes(), &respose)
		assert.NoError(err)
		assert.Equal("A011", respose.Member_id)
	})

	t.Run("Test_GetMemberinfobyIDAPI", func(t *testing.T) {

		req, _ := http.NewRequest("GET", "/api/members/A001", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(http.StatusOK, w.Code)
		var respose models.Member

		err := json.Unmarshal(w.Body.Bytes(), &respose)
		assert.NoError(err)
		assert.Equal("A001", respose.Member_id)

		// testing - accessing data which doesnt exist

		req, _ = http.NewRequest("GET", "/api/members/A012", nil)
		w = httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(http.StatusInternalServerError, w.Code)
	})

	t.Run("Test_DeleteMemberAPI", func(t *testing.T) {

		memberID := "A011"
		phone := "123456789"

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/members?id=%s&phone=%s", memberID, phone), nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(http.StatusOK, w.Code)

		// deleting data which doesnt exist
		req, _ = http.NewRequest("DELETE", "/api/members?id=A011&phone=123456789", nil)
		w = httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(http.StatusInternalServerError, w.Code)
	})

	// Test for Books routes
	t.Run("Test_Creating_BookAPI", func(t *testing.T) {
		testBook := models.Book{
			Title:    "TestName",
			Author:   "Author1",
			Category: "Cat1",
			Count:    1,
		}
		testBookJson, _ := json.Marshal(testBook)

		req, _ := http.NewRequest("POST", "/api/books", strings.NewReader(string(testBookJson)))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(http.StatusOK, w.Code)

		var response models.Book
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(err)
		assert.Equal(37, response.Book_id)
		assert.Equal("TestName", response.Title)
	})

	t.Run("Test_Filter_API", func(t *testing.T) {

		req, _ := http.NewRequest("GET", "/api/books?author=Author1", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(http.StatusOK, w.Code)

		var response []models.Book
		err := json.Unmarshal(w.Body.Bytes(), &response)

		assert.NoError(err)
		assert.Equal(1, len(response))
		assert.Equal("TestName", response[0].Title)

		// More tests
		tests := []struct {
			name          string
			author        string
			category      string
			prefix        string
			expectedCount int
		}{
			{"filter by title author", "J.R.R. Tolkien", "", "", 2},
			{"filter by title category", "", "Fantasy", "", 3},
			{"filter by title prefix", "", "", "Brave", 1},
			{"filter by title author and category", "J.R.R. Tolkien", "Fantasy", "", 2},
			{"filter by All", "Fyodor Dostoevsky", "Psychological", "Crime", 1},
			{"filter by nothing - gets all books", "", "", "", 37},
		}

		for _, test := range tests {
			var params []string
			if test.author != "" {
				params = append(params, fmt.Sprintf("author=%s", test.author))
			}
			if test.category != "" {
				params = append(params, fmt.Sprintf("category=%s", test.category))
			}
			if test.prefix != "" {
				params = append(params, fmt.Sprintf("prefix=%s", test.prefix))
			}
			url := "/api/books?" + strings.Join(params, "&")
			req, _ := http.NewRequest("GET", url, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)
			assert.Equal(http.StatusOK, w.Code)

			var response []models.Book
			err := json.Unmarshal(w.Body.Bytes(), &response)

			assert.NoError(err)
			assert.Equal(test.expectedCount, len(response), test.name)

		}
	})

	t.Run("Test_DeleteBookAPI", func(t *testing.T) {
		bookID := 37

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/books/%d", bookID), nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(http.StatusOK, w.Code)
		assert.NotNil(w.Body.String())

		//  Delete Book which doesnt exists
		bookID = 37

		req, _ = http.NewRequest("DELETE", fmt.Sprintf("/api/books/%d", bookID), nil)
		w = httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(http.StatusInternalServerError, w.Code)
	})

	// Test for Trasnaction Routes
	t.Run("Test_BorrowBook_API", func(t *testing.T) {
		request := struct {
			Memberid string `json:"member_id"`
			Bookid   int    `json:"book_id"`
		}{
			Memberid: "A001",
			Bookid:   1,
		}
		requestJson, _ := json.Marshal(&request)
		req, _ := http.NewRequest("POST", "/api/borrow", bytes.NewBuffer(requestJson))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(http.StatusOK, w.Code, w.Body.String())
		var response *models.Transaction

		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(err)
		assert.Equal(request.Memberid, response.Member.Member_id)

		// borrowing book from a Invalid Member
		request = struct {
			Memberid string `json:"member_id"`
			Bookid   int    `json:"book_id"`
		}{
			Memberid: "A011",
			Bookid:   1,
		}
		requestJson, _ = json.Marshal(&request)
		req, _ = http.NewRequest("POST", "/api/borrow", bytes.NewBuffer(requestJson))
		w = httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(http.StatusInternalServerError, w.Code)
	})

}
