package handlers

import (
	"github.com/askrishna98/library_/docs"
	"github.com/askrishna98/library_/loaddata"
	"github.com/askrishna98/library_/models"
	service "github.com/askrishna98/library_/services"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func Handlers(router *gin.Engine) {

	// initialize all

	MockDB := models.GetMockDBInstance()
	IdGenerator := service.InitalizeIDGenerator()

	MemberServices := service.GetInstanceOfMemberService(MockDB, IdGenerator)
	BookServices := service.GetInstanceOfBookService(MockDB, IdGenerator)
	TransactionServices := service.GetInstanceOfTransactionService(MockDB, IdGenerator, MemberServices, BookServices)

	//loading test data
	loaddata.LoadData(MockDB, IdGenerator, MemberServices, BookServices)

	router.GET("/", Greet)

	// Swagger BasePath
	docs.SwaggerInfo.BasePath = "/api"

	Group := router.Group("/api")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// member routes
	Group.POST("/members", CreateNewMember(MemberServices))
	Group.GET("/members/:id", GetUserByID(MemberServices))
	Group.DELETE("/members", DeleteMemberByID(MemberServices))

	// bookRoutes
	Group.POST("/books", CreateNewBook(BookServices))
	Group.DELETE("/books/:id", DeleteBookByID(BookServices))
	Group.GET("/books", Filter(BookServices))

	// transaction routes
	Group.POST("/borrow", BorrowBook(TransactionServices))
	Group.GET("/borrow/:id", GetListBooksByMemberID(TransactionServices))
	Group.PATCH("/return", ReturnBook(TransactionServices))
}

func StartApp() {

	router := gin.Default()

	Handlers(router)

	router.Run()
}
