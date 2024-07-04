package main

import (
	"github.com/askrishna98/library_/models"
	service "github.com/askrishna98/library_/services"
)

func main() {
	// should initiate DB first
	DB := models.GetMockDBInstance()
	newMember := models.Member{
		Member_id: "101",
		Name:      "ASWIN",
		Email:     "asdasd@gmail.com",
		Phone:     "3434134231",
	}

	newBook := models.Book{
		Title:    "test book",
		Category: "horro",
		Count:    2,
		Author:   "ram",
	}
	memberServices := service.GetInstanceOfMemberService(DB)
	bookServices := service.GetInstanceOfBookService(DB)
	memberServices.CreateMember(newMember)
	bookServices.CreateBook(newBook)

	Trans := service.GetInstanceOfTransactionService(DB, memberServices, bookServices)
	Trans.BorrowBook("101", 0)

	// id := service.InitalizeGenerator()
	// fmt.Println(id.Generate(), id.Generate())

}
