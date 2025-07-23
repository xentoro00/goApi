package members

import (
	"time"

	"github.com/google/uuid"
)

type Members struct {
	ID        int8      `json:"id,omitempty"`
	RoomID    uuid.UUID `json:"room_id"`
	MemberID  uuid.UUID `json:"member_id"`
	CreatedAt time.Time `json:"created_at"`
}
type MemberIDResult struct {
	MemberID uuid.UUID `json:"member_id"`
}
type RoomIDResult struct {
	RoomID uuid.UUID `json:"room_id"`
}
type CreateMemberRequest struct {
	RoomID uuid.UUID `json:"room_id" binding:"required"`
}
type GetMemberByIdRequest struct {
	ID uuid.UUID `json:"id" binding:"required"`
}
type GetMembersByRoomIdRequest struct {
	RoomID uuid.UUID `json:"room_id" binding:"required"`
}
type GetRoomsByMemberIdRequest struct {
	MemberID uuid.UUID `json:"member_id" binding:"required"`
}
type RoomsListResponse struct {
	Rooms []uuid.UUID `json:"rooms"`
}
