package rooms

import "github.com/google/uuid"

type Service interface {
	CreateRoom(req CreateRoomRequest, userID uuid.UUID) (*Room, error)
	GetAllRooms() ([]Room, error)
	GetById(id string) (*Room, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateRoom(req CreateRoomRequest, userID uuid.UUID) (*Room, error) {
	roomToCreate := Room{
		Name:      req.Name,
		CreatedBy: userID,
	}

	createdRoom, err := s.repo.Create(roomToCreate)
	if err != nil {
		return nil, err
	}

	return createdRoom, nil
}
func (s *service) GetAllRooms() ([]Room, error) {
	rooms, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
func (s *service) GetById(id string) (*Room, error) {
	return s.repo.GetByID(id)
}
