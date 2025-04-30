package handlers

import (
	"github.com/gin-gonic/gin"
	"go-rest-project/internal/models"
	"go-rest-project/internal/services"
	"strconv"
)

type HabitListHandler struct {
	habitListService services.HabitListService
}

func NewHabitListHandler(habitListService services.HabitListService) *HabitListHandler {
	return &HabitListHandler{habitListService: habitListService}
}

func (h *HabitListHandler) CreateHabitList(c *gin.Context) {
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

	var habitList models.HabitList
	if err := c.BindJSON(&habitList); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	habitList.UserID = userID
	if err := h.habitListService.CreateHabitList(&habitList); err != nil {
		switch err.Error() {
		case "habit list not found":
			c.JSON(404, gin.H{"error": err.Error()})
			return
		case "habit list with this name already exists":
			c.JSON(409, gin.H{"error": err.Error()})
			return
		default:
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(201, habitList)
}

func (h *HabitListHandler) GetHabitList(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil || id == 0 {
		c.JSON(400, gin.H{"error": "Invalid habit list ID"})
		return
	}

	habitList, err := h.habitListService.GetHabitListByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Habit list not found"})
		return
	}

	c.JSON(200, habitList)
}

func (h *HabitListHandler) GetUserHabitLists(c *gin.Context) {
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

	habitLists, err := h.habitListService.GetHabitListsByUserID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, habitLists)
}

func (h *HabitListHandler) UpdateHabitList(c *gin.Context) {
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

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil || id == 0 {
		c.JSON(400, gin.H{"error": "Invalid habit list ID"})
		return
	}

	var habitList models.HabitList
	if err := c.BindJSON(&habitList); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Проверка совпадения ID
	if habitList.ID != 0 && habitList.ID != uint(id) {
		c.JSON(400, gin.H{"error": "ID in body does not match ID in URL"})
		return
	}
	habitList.ID = uint(id)
	habitList.UserID = userID // Подставляем userID из контекста
	if err := h.habitListService.UpdateHabitList(&habitList); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, habitList)
}

func (h *HabitListHandler) DeleteHabitList(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil || id == 0 {
		c.JSON(400, gin.H{"error": "Invalid habit list ID"})
		return
	}

	if err := h.habitListService.DeleteHabitList(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Habit list deleted successfully"})
}
