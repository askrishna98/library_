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
	UniqPhones  sync.Map
}

func GetInstanceOfMemberService(DBInstance *models.MockDB, IdGeneratorInstance *IdGenerator) *MemberService {
	return &MemberService{
		DB:          DBInstance,
		IdGenerator: IdGeneratorInstance,
	}
}

// need to generate ID
// To add new member to Storage
func (m *MemberService) CreateMember(newMember *models.Member) error {

	if newMember.Name == "" || newMember.Phone == "" {
		return errors.New("name or phone field should not be empty")
	}
	phoneNum := newMember.Phone
	if _, ok := m.UniqPhones.Load(phoneNum); ok {
		return errors.New("phone number already exists")
	}

	newMember.Member_id = m.IdGenerator.GenerateMemberID()
	newMember.Date = time.Now().Format("02-01-2006")

	// Details of New member adds to DB
	m.DB.Members.Store(newMember.Member_id, newMember)
	// adds Phone number to uniqNumber
	m.UniqPhones.Store(phoneNum, newMember.Member_id)

	return nil
}

// to Delete member from Storage
func (m *MemberService) DeleteMember(memberId, phoneNumber string) error {

	value, ok := m.UniqPhones.Load(phoneNumber)
	// phone number do not exist
	if !ok {
		return errors.New("phonenumber do not exist")
	}
	// type assert
	memberID := value.(string)

	// No match
	if memberID != memberId {
		return errors.New("MemberID phoneNumber do not match")
	}
	value, err := m.GetMemberById(memberID)
	if err != nil {
		return err
	}
	member := value.(*models.Member)

	m.UniqPhones.Delete(phoneNumber)
	m.DB.Members.Delete(member.Member_id)

	return nil
}

// to update info
func (m *MemberService) UpdateMemberInfo(memberID string, Updateinfo models.MemberRequest) (models.Member, error) {

	newData := models.Member{}

	member, err := m.GetMemberById(memberID)

	if err != nil {
		return models.Member{}, err
	}

	newData = *member

	if Updateinfo.Name != "" {
		newData.Name = Updateinfo.Name
	}
	if Updateinfo.Email != "" {
		newData.Email = Updateinfo.Email
	}
	if Updateinfo.Phone != "" {
		m.UniqPhones.Delete(member.Phone)
		newData.Phone = Updateinfo.Phone
		m.UniqPhones.Store(newData.Phone, newData.Member_id)

	}

	m.DB.Members.Store(newData.Member_id, &newData)
	return newData, err
}

// to check whether memberID is Valid
func (m *MemberService) GetMemberById(memberID string) (*models.Member, error) {

	value, ok := m.DB.Members.Load(memberID)
	if ok {
		member := value.(*models.Member)
		return member, nil
	}
	return nil, errors.New("member not found")
}
