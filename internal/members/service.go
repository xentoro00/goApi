package members

import "github.com/google/uuid"

type Service interface {
	CreateMember(req CreateMemberRequest, userID uuid.UUID) (*Members, error)
	GetAllMembers() ([]Members, error)
	GetById(id string) (*Members, error)
	GetMembersByRoomID(req GetMembersByRoomIdRequest) ([]uuid.UUID, error)
	GetRoomsByMemberID(req GetRoomsByMemberIdRequest) ([]uuid.UUID, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateMember(req CreateMemberRequest, userID uuid.UUID) (*Members, error) {
	memberToCreate := Members{
		RoomID:   req.RoomID,
		MemberID: userID,
	}

	createdMember, err := s.repo.Create(memberToCreate)
	if err != nil {
		return nil, err
	}

	return createdMember, nil
}
func (s *service) GetAllMembers() ([]Members, error) {
	members, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return members, nil
}
func (s *service) GetById(id string) (*Members, error) {
	return s.repo.GetByID(id)
}
func (s *service) GetMembersByRoomID(req GetMembersByRoomIdRequest) ([]uuid.UUID, error) {
	members, err := s.repo.GetMembersByRoomID(req.RoomID)
	if err != nil {
		return nil, err
	}
	return members, nil
}
func (s *service) GetRoomsByMemberID(req GetRoomsByMemberIdRequest) ([]uuid.UUID, error) {
	rooms, err := s.repo.GetRoomsByMemberID(req.MemberID)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
