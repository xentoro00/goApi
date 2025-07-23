package members

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supabase-community/auth-go/types"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

// RegisterRoutes registers the member routes to the router.
func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/members", h.createMember)
	rg.GET("/members", h.getAllMembers)
	rg.GET("/members/:id", h.getMemberById)
	rg.POST("/getByRoom", h.GetMembersByRoomID)
	rg.POST("/getRoomsByMember", h.GetRoomsByMemberID)
}

// createMember handles the creation of a new member.
// It binds the JSON request body to the CreateMemberRequest struct.
func (h *Handler) createMember(c *gin.Context) {
	var req CreateMemberRequest

	// Call ShouldBindJSON only ONCE and store the error.
	// The request body is a stream and can only be read once.
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}

	// Assert the correct type: *types.UserResponse
	userData, ok := user.(*types.UserResponse)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user data in context"})
		return
	}

	// The UserResponse struct contains the ID we need
	log.Println("Creating member with request:", userData.ID)
	createdMember, err := h.service.CreateMember(req, userData.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create member: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdMember)
}

// getAllMembers handles GET /members requests
func (h *Handler) getAllMembers(c *gin.Context) {
	members, err := h.service.GetAllMembers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch members: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, members)
}

// getMemberById handles GET /members/:id requests
func (h *Handler) getMemberById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "member ID is required"})
		return
	}

	member, err := h.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch member: " + err.Error()})
		return
	}

	if member == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "member not found"})
		return
	}

	c.JSON(http.StatusOK, member)
}
func (h *Handler) GetMembersByRoomID(c *gin.Context) {
	var req GetMembersByRoomIdRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	members, err := h.service.GetMembersByRoomID(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch members: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"members": members})
}

func (h *Handler) GetRoomsByMemberID(c *gin.Context) {
	var req GetRoomsByMemberIdRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rooms, err := h.service.GetRoomsByMemberID(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch rooms: " + err.Error()})
		return
	}
	response := RoomsListResponse{Rooms: rooms}
	c.JSON(http.StatusOK, response)
}
