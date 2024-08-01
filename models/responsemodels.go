package models

type BookResponse struct {
	Book_id  int
	Title    string
	Author   string
	Category string
}

type BookTransactionResponse struct {
	Member struct {
		MemberID string
		Name     string
	}
	Books []struct {
		BookID       int
		Title        string
		Borrowdate   string
		DueDate      string
		ErrorMessage string
	}
}

type BookReturnTransactionResponse struct {
	Member struct {
		MemberID string
		Name     string
	}
	Books []struct {
		BookID        int
		Title         string
		Borrowdate    string
		DueDate       string
		Returned_Date string
		Penalty       int
		ErrorMessage  string
	}
}

type UpcomingReturnsResponse struct {
	MemberID   string
	MemberName string
	BookID     int
	Title      string
	BorrowDate string
	DueDate    string
}

// standard error response
type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessMessageResponse struct {
	Message string `json:"message"`
}
