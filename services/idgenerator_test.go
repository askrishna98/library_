package service

import "testing"

// Test for initialization
func TestInitalizeIDGenerator(t *testing.T) {
	idGenerator := InitalizeIDGenerator()
	if idGenerator == nil {
		t.Fatalf("Expected no nil object")
	}
	if idGenerator.num != 1 || idGenerator.alpha != "A" || idGenerator.bookId != 1 || idGenerator.transactionId != 1 {
		t.Fatalf("Initializaton failed : %+v", idGenerator)
	}
}

func TestGenerateBookId(t *testing.T) {
	idGenerator := InitalizeIDGenerator()
	bookid := idGenerator.GenerateBookId()
	if bookid != 1 {
		t.Fatalf("expected bookid to be 1 but got %d", bookid)
	}
	if idGenerator.bookId != 2 {
		t.Fatalf("expected next bookid to be 2 but got %d", idGenerator.bookId)
	}
}

func TestGenerateMemberID(t *testing.T) {
	IdGenerator := InitalizeIDGenerator()
	memberID := IdGenerator.GenerateMemberID()
	expectedID := "A001"
	if memberID != expectedID {
		t.Fatalf("expected MemberID to be %s but got %s", expectedID, memberID)
	}
}

func TestInitalizeNext(t *testing.T) {
	IdGenerator := &IdGenerator{
		alpha: "A",
		num:   999,
	}
	IdGenerator.InitalizeNext()
	expectedAlpha := "B"
	expectedNum := 1
	if IdGenerator.alpha != expectedAlpha || IdGenerator.num != expectedNum {
		t.Fatalf("Expected alpha %s and expected num %d, but got alpha %s and num %d", expectedAlpha, expectedNum, IdGenerator.alpha, IdGenerator.num)
	}
}

func TestNextAlpha(t *testing.T) {
	test := []struct {
		input    string
		expected string
	}{
		{"A", "B"},
		{"Z", "AA"},
		{"AA", "AB"},
		{"AZ", "BA"},
		{"ZZ", "AAA"},
		{"ZZZ", "AAAA"},
	}

	for _, test := range test {
		output := NextAlpha(test.input)
		if output != test.expected {
			t.Fatalf("for input %s expected is %s but got %s", test.input, test.expected, output)
		}
	}
}
