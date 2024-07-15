package handlers

import (
	"fmt"
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

func CreateNewMember(Mservice *service.MemberService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newMember models.Member
		if err := c.ShouldBindJSON(&newMember); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		Mservice.CreateMember(&newMember)
		c.JSON(http.StatusOK, newMember)
	}
}

// good to add phone number as well (use C.Query[para_name])
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

func Filter(Bservice *service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		author := c.Query("author")
		category := c.Query("category")
		prefix := c.Query("prefix")

		filtered_slice := Bservice.Filter(author, category, prefix)

		c.JSON(http.StatusOK, filtered_slice)
	}
}

func BorrowBook(Tservice *service.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Memberid string `json:"member_id"`
			Bookid   int    `json:"book_id"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if newTrans, err := Tservice.BorrowBook(request.Memberid, request.Bookid); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, *newTrans)
		}
	}
}

func ReturnBook(Tservice *service.TransactionService) gin.HandlerFunc {

	type ResponseWithPenalty struct {
		models.Transaction
		Penalty int
	}

	return func(c *gin.Context) {
		var request struct {
			Memberid string `json:"member_id"`
			Bookid   int    `json:"book_id"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if Trans, penalty, err := Tservice.ReturnBook(request.Memberid, request.Bookid); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		} else {

			c.JSON(http.StatusOK, ResponseWithPenalty{
				Transaction: *Trans,
				Penalty:     penalty,
			})
			return
		}

	}
}

func GetListBooksByMemberID(Tservice *service.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var books []models.Book

		memberID := c.Param("id")
		books, err := Tservice.GetBooksBorrowedByMember(memberID)
		fmt.Println(memberID)

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
