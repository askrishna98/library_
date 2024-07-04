package models

// list of all members
// var Members []Member = []Member{
// 	{Member_id: 1, Name: "Alice", Email: "alice@example.com", Phone: "123-456-7890", Date: "2023-01-01"},
// 	{Member_id: 2, Name: "Bob", Email: "bob@example.com", Phone: "123-456-7891", Date: "2023-02-01"},
// 	{Member_id: 3, Name: "Charlie", Email: "charlie@example.com", Phone: "123-456-7892", Date: "2023-03-01"},
// }

// MockDB
type MockDB struct {
	Members          []Member
	Books            []Book
	BookTransactions []*Transaction
}

//to get instance

func GetMockDBInstance() *MockDB {
	var instance MockDB
	return &instance
}
