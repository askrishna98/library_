package models

import "sync"

// MockDB
type MockDB struct {
	Books            []*Book
	BookTransactions []*Transaction
	Members          sync.Map
}

//to get instance

func GetMockDBInstance() *MockDB {
	var instance MockDB
	return &instance
}
