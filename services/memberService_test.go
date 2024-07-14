package service

import (
	"testing"

	"github.com/askrishna98/library_/models"
)

func TestCreateMember(t *testing.T) {
	mockDB := models.GetMockDBInstance()
	IdGenerator := InitalizeIDGenerator()
	MemberService := GetInstanceOfMemberService(mockDB, IdGenerator)

	testMember := &models.Member{
		Name:  "test1",
		Email: "test@gmail.com",
		Phone: "938849132112",
	}
	if err := MemberService.CreateMember(testMember); err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
	if len(mockDB.Members) != 1 {
		t.Fatalf("Expected number of memebers = 1 , got %d", len(mockDB.Members))
	}
	if testMember.Member_id != "A001" {
		t.Fatalf("expected memberid A001, but got %s", testMember.Member_id)
	}

}

func TestDeleteMember(t *testing.T) {
	mockDB := &models.MockDB{
		Members: []*models.Member{
			{
				Member_id: "A001",
				Name:      "testName",
				Phone:     "123456789",
			},
		},
	}
	IdGenerator := InitalizeIDGenerator()
	MemberService := GetInstanceOfMemberService(mockDB, IdGenerator)

	if err := MemberService.DeleteMember("A001", "123456789"); err != nil {
		t.Fatalf("expected no error but got %v", err)
	}
	if len(mockDB.Members) != 0 {
		t.Fatalf("Expected length of members slice to be 0 4, but got %d", len(mockDB.Members))
	}
}

func TestGetMemberById(t *testing.T) {
	mockDB := &models.MockDB{
		Members: []*models.Member{
			{
				Member_id: "A001",
				Name:      "testName",
				Phone:     "123456789",
			},
		},
	}
	IdGenerator := InitalizeIDGenerator()
	MemberService := GetInstanceOfMemberService(mockDB, IdGenerator)

	member, err := MemberService.GetMemberById("A001")
	if err != nil {
		t.Fatalf("expected No error but got error '%v'", err)
	}
	if member.Member_id != "A001" {
		t.Fatalf("expected meberid to be A001 but got %s", member.Member_id)
	}

}
