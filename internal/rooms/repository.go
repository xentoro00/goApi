package rooms

import (
	supa "github.com/nedpals/supabase-go"
)

// Repository defines the interface for room database operations.
type Repository interface {
	Create(room Room) (*Room, error)
	GetAll() ([]Room, error)
	GetByID(id string) (*Room, error)
}

type repository struct {
	db *supa.Client
}

// NewRepository creates a new room repository.
func NewRepository(db *supa.Client) Repository {
	return &repository{db: db}
}

func (r *repository) Create(room Room) (*Room, error) {
	var results []Room

	insertData := map[string]string{
		"name":       room.Name,
		"created_by": room.CreatedBy.String(),
	}

	err := r.db.DB.From("rooms").Insert(insertData).Execute(&results)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, nil
	}

	return &results[0], nil
}

func (r *repository) GetAll() ([]Room, error) {
	var results []Room

	err := r.db.DB.From("rooms").Select("*").Execute(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *repository) GetByID(id string) (*Room, error) {
	// FIX: The result from Execute must be a slice.
	var results []Room

	err := r.db.DB.From("rooms").Select("*").Eq("id", id).Execute(&results)
	if err != nil {
		return nil, err
	}

	// If the slice is empty, the room was not found.
	if len(results) == 0 {
		return nil, nil
	}

	// Otherwise, return the first element of the slice.
	return &results[0], nil
}
