package handlers

import (
	"github.com/gin-gonic/gin"
	"go-rest-project/internal/models"
	"go-rest-project/internal/services"
	"strconv"
)

type HabitHandler struct {
	habitService services.HabitService
}

func NewHabitHandler(habitService services.HabitService) *HabitHandler {
	return &HabitHandler{habitService: habitService}
}

func (h *HabitHandler) CreateHabit(c *gin.Context) {
	// userID, username, role уже извлекаются, но HabitListID должен быть в теле или параметрах

	okID := false
okUsername := false
okRole := false
_, okID = c.Get("userID")
_, okUsername = c.Get("username")
_, okRole = c.Get("role")
if !okID || !okUsername || !okRole {
	c.JSON(401, gin.H{"error": "User context (userID, username, role) not found in context"})
	return
}

	var habit models.Habit
	if err := c.BindJSON(&habit); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	// Если в теле нет habit_list_id, но есть в query, подставим
	if habit.HabitListID == 0 {
		listIDStr := c.Query("list_id")
		if listIDStr != "" {
			listID, err := strconv.ParseUint(listIDStr, 10, 32)
			if err == nil && listID > 0 {
				habit.HabitListID = uint(listID)
			}
		}
	}

	if err := h.habitService.CreateHabit(&habit); err != nil {
		switch err.Error() {
		case "referenced habit list does not exist":
			c.JSON(404, gin.H{"error": err.Error()})
			return
		default:
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(201, habit)
}

func (h *HabitHandler) GetHabit(c *gin.Context) {
	okID := false
okUsername := false
okRole := false
_, okID = c.Get("userID")
_, okUsername = c.Get("username")
_, okRole = c.Get("role")
if !okID || !okUsername || !okRole {
	c.JSON(401, gin.H{"error": "User context (userID, username, role) not found in context"})
	return
}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil || id == 0 {
		c.JSON(400, gin.H{"error": "Invalid habit ID"})
		return
	}

	habit, err := h.habitService.GetHabitByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Habit not found"})
		return
	}

	c.JSON(200, habit)
}

func (h *HabitHandler) GetHabitsByList(c *gin.Context) {
	okID := false
okUsername := false
okRole := false
_, okID = c.Get("userID")
_, okUsername = c.Get("username")
_, okRole = c.Get("role")
if !okID || !okUsername || !okRole {
	c.JSON(401, gin.H{"error": "User context (userID, username, role) not found in context"})
	return
}
	listIDStr := c.Query("list_id")
	if listIDStr == "" {
		c.JSON(400, gin.H{"error": "Missing list_id parameter"})
		return
	}
	listID, err := strconv.ParseUint(listIDStr, 10, 32)
	if err != nil || listID == 0 {
		c.JSON(400, gin.H{"error": "Invalid habit list ID"})
		return
	}

	habits, err := h.habitService.GetHabitsByHabitListID(uint(listID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, habits)
}

func (h *HabitHandler) UpdateHabit(c *gin.Context) {
	// userID, username, role уже извлекаются, но HabitListID должен быть подставлен явно, если есть

	okID := false
okUsername := false
okRole := false
_, okID = c.Get("userID")
_, okUsername = c.Get("username")
_, okRole = c.Get("role")
if !okID || !okUsername || !okRole {
	c.JSON(401, gin.H{"error": "User context (userID, username, role) not found in context"})
	return
}
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid habit ID"})
		return
	}

	var habit models.Habit
	if err := c.BindJSON(&habit); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	// Если в теле нет habit_list_id, но есть в query, подставим
	if habit.HabitListID == 0 {
		listIDStr := c.Query("list_id")
		if listIDStr != "" {
			listID, err := strconv.ParseUint(listIDStr, 10, 32)
			if err == nil && listID > 0 {
				habit.HabitListID = uint(listID)
			}
		}
	}

	// Проверка совпадения ID
	if habit.ID != 0 && habit.ID != uint(id) {
		c.JSON(400, gin.H{"error": "ID in body does not match ID in URL"})
		return
	}
	habit.ID = uint(id)
	if err := h.habitService.UpdateHabit(&habit); err != nil {
		switch err.Error() {
		case "habit not found":
			c.JSON(404, gin.H{"error": err.Error()})
			return
		case "habit with this name already exists":
			c.JSON(409, gin.H{"error": err.Error()})
			return
		default:
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, habit)
}

func (h *HabitHandler) DeleteHabit(c *gin.Context) {
	okID := false
okUsername := false
okRole := false
_, okID = c.Get("userID")
_, okUsername = c.Get("username")
_, okRole = c.Get("role")
if !okID || !okUsername || !okRole {
	c.JSON(401, gin.H{"error": "User context (userID, username, role) not found in context"})
	return
}
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil || id == 0 {
		c.JSON(400, gin.H{"error": "Invalid habit ID"})
		return
	}

	if err := h.habitService.DeleteHabit(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Habit deleted successfully"})
}
