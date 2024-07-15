package service

import (
	"errors"
	"sync"
	"time"

	"github.com/askrishna98/library_/models"
)

type MemberService struct {
	DB          *models.MockDB
	IdGenerator *IdGenerator
	mutex       *sync.Mutex
}

func GetInstanceOfMemberService(DBInstance *models.MockDB, IdGeneratorInstance *IdGenerator) *MemberService {
	return &MemberService{
		DB:          DBInstance,
		IdGenerator: IdGeneratorInstance,
		mutex:       &sync.Mutex{},
	}
}

// need to generate ID
// To add new member to Storage
func (m *MemberService) CreateMember(newMember *models.Member) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if newMember.Name == "" {
		return errors.New("name field should not be empty")
	}

	newMember.Member_id = m.IdGenerator.GenerateMemberID()
	newMember.Date = time.Now().Format("02-01-2006")
	m.DB.Members = append(m.DB.Members, newMember)
	return nil
}

// to Delete member from Storage
func (m *MemberService) DeleteMember(memberId, phoneNumber string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for i, member := range m.DB.Members {
		if memberId == member.Member_id {
			if member.Phone == phoneNumber {
				m.DB.Members = append(m.DB.Members[:i], m.DB.Members[i+1:]...)
			} else {
				return errors.New("phonenumber do not match")
			}
			return nil
		}
	}
	return errors.New("member not found")
}

// to check whether memberID is Valid
func (m *MemberService) GetMemberById(memberID string) (*models.Member, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for _, member := range m.DB.Members {
		if memberID == member.Member_id {
			return member, nil
		}
	}
	return nil, errors.New("member not found")
}
