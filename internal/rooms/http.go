package rooms

import (
	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
	"net/http"

	"github.com/supabase-community/auth-go/types"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

// RegisterRoutes registers the room routes to the router.
func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/rooms", h.createRoom)
	rg.GET("/rooms", h.getAllRooms)
	rg.GET("/rooms/:id", h.getRoomById)
}

func (h *Handler) createRoom(c *gin.Context) {
	var req CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}

	// FIX: Assert the correct type: *types.UserResponse
	userData, ok := user.(*types.UserResponse)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user data in context"})
		return
	}

	// The UserResponse struct contains the ID we need
	createdRoom, err := h.service.CreateRoom(req, userData.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create room" + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdRoom)
}

// getAllRooms handles GET /rooms requests
func (h *Handler) getAllRooms(c *gin.Context) {
	rooms, err := h.service.GetAllRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch rooms: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, rooms)
}

func (h *Handler) getRoomById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room ID is required"})
		return
	}

	room, err := h.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch room: " + err.Error()})
		return
	}

	if room == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "room not found"})
		return
	}

	c.JSON(http.StatusOK, room)
}
