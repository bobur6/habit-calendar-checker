package handlers

import (
	"github.com/gin-gonic/gin"
	"go-rest-project/internal/auth"
	"go-rest-project/internal/models"
	"go-rest-project/internal/services"
	"strings"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateProfileRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := h.userService.Register(user); err != nil {
		switch err {
		case services.ErrUsernameExists, services.ErrEmailExists:
			c.JSON(409, gin.H{"error": err.Error()})
			return
		default:
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(201, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if strings.TrimSpace(req.Username) == "" || strings.TrimSpace(req.Password) == "" {
		c.JSON(400, gin.H{"error": "Username and password are required"})
		return
	}

	user, err := h.userService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := auth.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userIDVal, ok := c.Get("userID")
	if !ok {
		c.JSON(401, gin.H{"error": "User context (userID) not found in context"})
		return
	}
	userID, ok := userIDVal.(uint)
	if !ok {
		c.JSON(401, gin.H{"error": "userID in context is not uint"})
		return
	}

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userIDVal, ok := c.Get("userID")
	if !ok {
		c.JSON(401, gin.H{"error": "User context (userID) not found in context"})
		return
	}
	userID, ok := userIDVal.(uint)
	if !ok {
		c.JSON(401, gin.H{"error": "userID in context is not uint"})
		return
	}

	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	currentUser, err := h.userService.GetUserByID(userID)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if req.Username != "" {
		currentUser.Username = strings.TrimSpace(req.Username)
	}
	if req.Email != "" {
		currentUser.Email = strings.TrimSpace(req.Email)
	}
	if req.Password != "" {
		if len(req.Password) < 6 {
			c.JSON(400, gin.H{"error": "Password must be at least 6 characters long"})
			return
		}
		currentUser.Password = req.Password
	}

	if err := h.userService.UpdateUser(currentUser); err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			c.JSON(409, gin.H{"error": "Username or email already exists"})
			return
		}
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"id":       currentUser.ID,
		"username": currentUser.Username,
		"email":    currentUser.Email,
	})
}

func (h *UserHandler) DeleteProfile(c *gin.Context) {
	userIDVal, ok := c.Get("userID")
	if !ok {
		c.JSON(401, gin.H{"error": "User context (userID) not found in context"})
		return
	}
	userID, ok := userIDVal.(uint)
	if !ok {
		c.JSON(401, gin.H{"error": "userID in context is not uint"})
		return
	}

	if err := h.userService.DeleteUser(userID); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}
