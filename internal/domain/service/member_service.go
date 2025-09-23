package service

import "github.com/google/uuid"

type MemberServiceInterface interface {
	IsActive(id uuid.UUID) (bool, error)
}

type MemberService struct {
}

func NewMemberService() *MemberService {
	return &MemberService{}
}

func (m *MemberService) IsActive(id uuid.UUID) (bool, error) {
	return true, nil
}
