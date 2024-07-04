package service

import "fmt"

type IdGenerator struct {
	Num           int
	Alpha         string
	BookID        int
	TransactionID int
}

func InitalizeIDGenerator() *IdGenerator {
	return &IdGenerator{
		Num:           1,
		Alpha:         "A",
		BookID:        0,
		TransactionID: 0,
	}
}

// just to create bookID
func (i *IdGenerator) GenerateBookID() int {
	id := i.BookID

	i.BookID++
	return id
}

func (i *IdGenerator) GenerateTransactionID() int {
	id := i.TransactionID

	i.BookID++
	return id
}

func (i *IdGenerator) GenerateMemberID() string {
	ID := i.Alpha + fmt.Sprintf("%03d", i.Num)
	i.InitalizeNext()
	return ID

}

func (i *IdGenerator) InitalizeNext() {
	if i.Num == 999 {
		i.Num = 1
		i.Alpha = nextAlpha(i.Alpha)
	} else {
		i.Num++
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
