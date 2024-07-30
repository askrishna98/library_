package models

import (
	"sync"
)

type List struct {
	List []*interface{}
	mu   *sync.Mutex
}

func InitializeNewList() *List {
	return &List{
		List: []*interface{}{},
		mu:   &sync.Mutex{},
	}
}

func (l *List) AddNewItem(newItem any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.List = append(l.List, &newItem)
}

func (l *List) DeleteItem(index int) interface{} {
	l.mu.Lock()
	defer l.mu.Unlock()

	deletedItem := l.List[index]
	l.List = append(l.List[:index], l.List[index+1:]...)

	return deletedItem
}

func (l *List) GetItems() *[]*interface{} {
	l.mu.Lock()
	defer l.mu.Unlock()

	return &l.List
}

func (l *List) Len() int {
	l.mu.Lock()
	defer l.mu.Unlock()

	return len(l.List)
}
