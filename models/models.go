package models

import "sync"

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
	mu       sync.Mutex
	Book_id  int
	Title    string
	Author   string
	Category string
	Count    int
}

func (b *Book) IncrementCount() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.Count++
}

func (b *Book) DecrementCount() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.Count--
}
func (b *Book) UpdateCount(newVal int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Count = newVal
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
