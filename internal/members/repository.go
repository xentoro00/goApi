package members

import (
	"log"

	"github.com/google/uuid"
	supa "github.com/nedpals/supabase-go"
)

// Repository defines the interface for member database operations.
type Repository interface {
	Create(member Members) (*Members, error)
	GetAll() ([]Members, error)
	GetByID(id string) (*Members, error)
	GetMembersByRoomID(roomID uuid.UUID) ([]uuid.UUID, error)
	GetRoomsByMemberID(memberID uuid.UUID) ([]uuid.UUID, error)
}

type repository struct {
	db *supa.Client
}

// NewRepository creates a new room repository.
func NewRepository(db *supa.Client) Repository {
	return &repository{db: db}
}

func (r *repository) Create(member Members) (*Members, error) {
	var results []Members

	insertData := map[string]string{
		"room_id":   member.RoomID.String(),
		"member_id": member.MemberID.String(),
	}

	err := r.db.DB.From("members").Insert(insertData).Execute(&results)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, nil
	}

	return &results[0], nil
}

func (r *repository) GetAll() ([]Members, error) {
	var results []Members

	err := r.db.DB.From("members").Select("*").Execute(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *repository) GetByID(id string) (*Members, error) {
	// FIX: The result from Execute must be a slice.
	var results []Members

	err := r.db.DB.From("members").Select("*").Eq("id", id).Execute(&results)
	if err != nil {
		return nil, err
	}

	// If the slice is empty, the member was not found.
	if len(results) == 0 {
		return nil, nil
	}

	// Otherwise, return the first element of the slice.
	return &results[0], nil
}

func (r *repository) GetMembersByRoomID(roomID uuid.UUID) ([]uuid.UUID, error) {
	// This will hold the raw results from the database query.
	var queryResults []MemberIDResult
	log.Println("Fetching members for roomID:", roomID.String())

	// Execute the query and scan the results into the slice of structs.
	err := r.db.DB.From("members").Select("member_id").Eq("room_id", roomID.String()).Execute(&queryResults)
	if err != nil {
		return nil, err
	}

	// Now, extract the UUIDs from the structs into the final slice.
	finalResults := make([]uuid.UUID, len(queryResults))
	for i, item := range queryResults {
		finalResults[i] = item.MemberID
	}

	return finalResults, nil
}

func (r *repository) GetRoomsByMemberID(memberID uuid.UUID) ([]uuid.UUID, error) {
	var queryResults []RoomIDResult

	err := r.db.DB.From("members").Select("room_id").Eq("member_id", memberID.String()).Execute(&queryResults)
	if err != nil {
		return nil, err
	}

	finalResults := make([]uuid.UUID, len(queryResults))
	for i, item := range queryResults {
		finalResults[i] = item.RoomID
	}

	return finalResults, nil

}
