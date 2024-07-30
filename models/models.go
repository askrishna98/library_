package models

// member struct
type Member struct {
	Member_id string `json: "member_id"`
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
