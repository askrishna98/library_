package service

import "fmt"

type IdGenerator struct {
	num           int
	alpha         string
	bookId        int
	transactionId int
}

func InitalizeIDGenerator() *IdGenerator {
	return &IdGenerator{
		num:           1,
		alpha:         "A",
		bookId:        0,
		transactionId: 0,
	}
}

// Generate bookId from 0 - ....(int value)
func (i *IdGenerator) GenerateBookId() int {
	id := i.bookId

	i.bookId++
	return id
}

// // Generate unique Id for every transaction from 0 - ....(int value)
func (i *IdGenerator) GenerateTransactionId() int {
	id := i.transactionId

	i.transactionId++
	return id
}

func (i *IdGenerator) GenerateMemberID() string {
	ID := i.alpha + fmt.Sprintf("%03d", i.num)
	i.InitalizeNext()
	return ID

}

func (i *IdGenerator) InitalizeNext() {
	if i.num == 999 {
		i.num = 1
		i.alpha = nextAlpha(i.alpha)
	} else {
		i.num++
	}
}

func nextAlpha(s string) string {
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
