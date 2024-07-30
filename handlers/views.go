package handlers

import (
	"net/http"
	"strconv"

	"github.com/askrishna98/library_/models"
	service "github.com/askrishna98/library_/services"
	"github.com/gin-gonic/gin"
)

func Greet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"greet": "HELLO WORLD",
	})
}

// @BasePath /api/v1

// CreateNewMember godoc
// @Summary Creates a new Member
// @Description Creates a new member, details should be passed in JSON. name and phone number is mandatory
// @Tags members
// @Accept json
// @Produce json
// @Param member body models.MemberRequest true "Member details"
// @Success 200 {object} models.Member "Member details"
// @Failure 400 {object} models.ErrorResponse "error message"
// @Router /members [post]
func CreateNewMember(Mservice *service.MemberService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newMember models.Member
		if err := c.ShouldBindJSON(&newMember); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := Mservice.CreateMember(&newMember); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, newMember)
	}
}

// good to add phone number as well (use C.Query[para_name])

// GetUserByID godoc
// @Summary Get Member by Member ID
// @Description Get details of a member by their ID
// @Tags members
// @Param id path string true "Member ID"
// @Success 200 {object} models.Member "Member details"
// @Failure 500 {object} models.ErrorResponse "error message"
// @Router /members/{id} [get]
func GetUserByID(Mservice *service.MemberService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		member, err := Mservice.GetMemberById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, member)

	}
}

// DeleteMemberByID godoc
// @Summary To delete a Member
// @Description Deletes the member using MemberID and PhoneNumber passed as query parameters
// @Tags members
// @Param id query string true "Member ID"
// @Param phone query string true "Phone Number"
// @Success 200 {object} models.SuccessMessageResponse "success message"
// @Failure 500 {object} models.ErrorResponse "error message"
// @Router /members [delete]
func DeleteMemberByID(Mservice *service.MemberService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		phone := c.Query("phone")
		err := Mservice.DeleteMember(id, phone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Deleted successfully",
		})
	}
}

// book Functionalities

// CreateNewBook godoc
// @Summary Creates a new Book
// @Description Creates a new Book, details should be passed in JSON.
// @Tags Books
// @Accept json
// @Produce json
// @Param book body models.BookRequest true "Book details"
// @Success 200 {object} models.Book "Book details"
// @Failure 400 {object} models.ErrorResponse "error message"
// @Router /books [post]
func CreateNewBook(Bservice *service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newBook models.Book
		if err := c.ShouldBindJSON(&newBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		err := Bservice.CreateBook(&newBook)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, newBook)

	}
}

// DeleteBookByID godoc
// @Summary  To Delete Book
// @Description Deletes a existing book By its ID
// @Tags Books
// @Param id path string true "BookID"
// @Success 200 {object} models.SuccessMessageResponse "success message"
// @Failure 500 {object} models.ErrorResponse "Error message"
// @Router /books/{id} [delete]
func DeleteBookByID(Bservice *service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		BookId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := Bservice.DeleteBook(BookId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"messasge": "Book Deleted successfully",
		})
	}
}

// Filter godoc
// @Summary Filters books
// @Description Filters all books by query parameters 'author','category','prefix'.all paramters are not mandatory. details of all books will be given if no paramters are provided.
// @Tags Books
// @Produce json
// @Param author query string false "author"
// @Param category query string false "category"
// @Param prefix query string false "prefix"
// @Success 200 {object} []models.Book "List of Books"
// @Failure 400 {object} models.ErrorResponse "error message"
// @Router /books [get]
func Filter(Bservice *service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		author := c.Query("author")
		category := c.Query("category")
		prefix := c.Query("prefix")

		filtered_slice := Bservice.Filter(author, category, prefix)

		c.JSON(http.StatusOK, filtered_slice)
	}
}

// Book Transactions functionalities

// BorrowBook godoc
// @Summary Creates a new Book Transaction
// @Description Creates a new Book Transaction, member_id and book_id should be passed in JSON.
// @Tags Book-Transactions
// @Accept json
// @Produce json
// @Param book body models.BookTransactionRequest true "Request for bookTransaction"
// @Success 200 {object} models.BookTransactionResponse "Book-Transaction details"
// @Failure 400 {object} models.ErrorResponse "error message"
// @Router /borrow [post]
func BorrowBook(Tservice *service.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {

		var request models.BookTransactionRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if newTrans, err := Tservice.BorrowBooks(request.Memberid, request.Bookids); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, *newTrans)
		}
	}
}

// ReturnBook godoc
// @Summary Updates the Book-transaction
// @Description Updates the book-transaction, returned_date and penalty is populated in the system, member_id and book_id should be passed in JSON.
// @Tags Book-Transactions
// @Accept json
// @Produce json
// @Param book body models.BookTransactionRequest true "Request for  returnbook"
// @Success 200 {object} models.BookReturnTransactionResponse "details of ReturnedBooks"
// @Failure 400 {object} models.ErrorResponse "error message"
// @Router /return [patch]
func ReturnBook(Tservice *service.TransactionService) gin.HandlerFunc {

	return func(c *gin.Context) {
		var request models.BookTransactionRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if newTrans, err := Tservice.ReturnBooks(request.Memberid, request.Bookids); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		} else {

			c.JSON(http.StatusOK, newTrans)
			return
		}

	}
}

// GetListBooksByMemberID godoc
// @Summary Gets list of all books borrowed By member
// @Description Gets list of all books books borrowed by memberID which are not returned yet.
// @Tags Book-Transactions
// @Produce json
// @Param id path string true "member_id"
// @Success 200 {object} []models.Book "details of books"
// @Failure 400 {object} models.ErrorResponse "error message"
// @Router /borrow/{id} [get]
func GetListBooksByMemberID(Tservice *service.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var books []models.Book

		memberID := c.Param("id")
		books, err := Tservice.GetBooksBorrowedByMember(memberID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		if len(books) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Member has not borrowed any books currently.",
			})
			return
		}
		c.JSON(http.StatusOK, books)
	}
}

// UpcomingReturns godoc
// @Summary Gets list of all Upcoming Returns
// @Description Gets list of all upcoming returns of books in timeframe. Expect date ("DD-MM-YYYY") as query paramter not mandatory, and lists all books which has due date before the date provided. All upcoming books will belisted if no date provided.
// @Tags Book-Transactions
// @Produce json
// @Param date query string false "date"
// @Success 200 {object} []models.UpcomingReturnsResponse "details of Upcomingbooks and member"
// @Failure 400 {object} models.ErrorResponse "error message"
// @Router /return [get]
func UpcomingReturns(Tservice *service.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {

		date := c.Query("date")
		res := Tservice.UpcomingReturns(date)
		if len(res) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": "No Books to be returned",
			})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}
