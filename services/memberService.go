package service

import (
	"errors"

	"github.com/askrishna98/library_/models"
)

type MemberService struct {
	DB *models.MockDB
}

func GetInstanceOfMemberService(DBInstance *models.MockDB) *MemberService {
	return &MemberService{
		DB: DBInstance,
	}
}

// need to generate ID
// To add new member to Storage
func (m MemberService) CreateMember(newMember models.Member) error {
	m.DB.Members = append(m.DB.Members, newMember)
	return nil
}

// to Delete member from Storage
func (m MemberService) DeleteMember(memberId string) error {
	for i, member := range m.DB.Members {
		if memberId == member.Member_id {
			m.DB.Members = append(m.DB.Members[:i], m.DB.Members[i+1:]...)
			return nil
		}
	}
	return errors.New("member not found")
}

// to check whether memberID is Valid
func (m MemberService) GetMemberById(memberID string) (*models.Member, error) {
	for _, member := range m.DB.Members {
		if memberID == member.Member_id {
			return &member, nil
		}
	}
	return nil, errors.New("member not found")
}
