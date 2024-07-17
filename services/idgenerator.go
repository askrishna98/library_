package service

import (
	"fmt"
	"sync"
)

type IdGenerator struct {
	num           int
	alpha         string
	bookId        int
	transactionId int
	mutex         *sync.Mutex
}

func InitalizeIDGenerator() *IdGenerator {
	return &IdGenerator{
		num:           1,
		alpha:         "A",
		bookId:        1,
		transactionId: 1,
		mutex:         &sync.Mutex{},
	}
}

// Generate bookId from 0 - ....(int value)
func (i *IdGenerator) GenerateBookId() int {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	id := i.bookId

	i.bookId++
	return id
}

// // Generate unique Id for every transaction from 0 - ....(int value)
func (i *IdGenerator) GenerateTransactionId() int {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	id := i.transactionId

	i.transactionId++
	return id
}

func (i *IdGenerator) GenerateMemberID() string {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	ID := i.alpha + fmt.Sprintf("%03d", i.num)
	i.InitalizeNext()
	return ID

}

func (i *IdGenerator) InitalizeNext() {

	if i.num == 999 {
		i.num = 1
		i.alpha = NextAlpha(i.alpha)
	} else {
		i.num++
	}
}

func NextAlpha(s string) string {

	runes := []rune(s)

	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] != 'Z' {
			runes[i]++
			return string(runes)
		} else {
			runes[i] = 'A'
		}
	}
	return "A" + string(runes)
}
