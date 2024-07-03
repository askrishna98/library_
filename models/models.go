package models

// member struct
type Member struct {
	Member_id string
	Name      string
	Email     string
	Phone     string
	Date      string
}

// Book Struct
type Book struct {
	Book_id  int
	Title    string
	Author   string
	Category string
	Count    int
}

// book Transaction
type Transaction struct {
	Borrow_id   int
	Member      *Member
	Book        *Book
	Borrow_date string
	Due_date    string
	Return_date string
}

// func CreateMember(newMember *Member) error
// func DeleteMember(member *Member) error
// func getMemberId(id string)

// func FindBookByAuthor(author string)
// func FindBookByCategory(category string)
// func FindBookByPrefix(prefix string)
// func CreateBook(newBook *Book)
// func DeleteBook(book *Book)

// func BorrowBook(member_id string, book_id int)
// func ReturnBook(member_id string, book_id int)
