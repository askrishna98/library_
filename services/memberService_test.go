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
	if _, ok := mockDB.Members.Load(testMember.Member_id); !ok {
		t.Fatalf("Expected MemberID - %s  present in DB", testMember.Member_id)
	}
	if testMember.Member_id != "A001" {
		t.Fatalf("expected memberid A001, but got %s", testMember.Member_id)
	}

}

func TestDeleteMember(t *testing.T) {
	mockDB := &models.MockDB{}
	newMember := &models.Member{
		Name:  "testName",
		Phone: "123456789",
	}
	IdGenerator := InitalizeIDGenerator()
	MemberService := GetInstanceOfMemberService(mockDB, IdGenerator)
	MemberService.CreateMember(newMember)

	if err := MemberService.DeleteMember("A001", "123456789"); err != nil {
		t.Fatalf("expected no error but got %v", err)
	}
	if _, ok := mockDB.Members.Load(newMember.Member_id); ok {
		t.Fatalf("Expected MemberID %s  not present in DB", newMember.Member_id)
	}
}

func TestGetMemberById(t *testing.T) {
	mockDB := &models.MockDB{}
	newMember := &models.Member{
		Name:  "testName",
		Phone: "123456789",
	}
	IdGenerator := InitalizeIDGenerator()
	MemberService := GetInstanceOfMemberService(mockDB, IdGenerator)
	MemberService.CreateMember(newMember)

	member, err := MemberService.GetMemberById("A001")
	if err != nil {
		t.Fatalf("expected No error but got error '%v'", err)
	}
	if member.Member_id != "A001" {
		t.Fatalf("expected meberid to be A001 but got %s", member.Member_id)
	}

}
