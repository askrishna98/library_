package models

// MockDB
type MockDB struct {
	Members          []*Member
	Books            []*Book
	BookTransactions []*Transaction
}

//to get instance

func GetMockDBInstance() *MockDB {
	var instance MockDB
	return &instance
}
