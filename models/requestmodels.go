package models

// transaction request body
type BookTransactionRequest struct {
	Memberid string `json:"member_id"`
	Bookids  []int  `json:"book_ids"`
}

// Member Request
type MemberRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// BookRequest
type BookRequest struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Category string `json:"category"`
	Count    int    `json:"count"`
}
