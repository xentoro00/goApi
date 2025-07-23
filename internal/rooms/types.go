package rooms

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name"`
	CreatedBy uuid.UUID `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateRoomRequest struct {
	Name string `json:"name" binding:"required"`
}
type GetRoomByIdRequest struct {
	Id uuid.UUID `json:"id" binding:"required"`
}
