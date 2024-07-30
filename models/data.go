package models

import "sync"

// MockDB
type MockDB struct {
	// Books            []*Book
	BookTransactions []*Transaction
	Members          sync.Map
	Books            *List
}

//to get instance

func GetMockDBInstance() *MockDB {
	instance := MockDB{
		Books: InitializeNewList(),
	}
	return &instance
}
