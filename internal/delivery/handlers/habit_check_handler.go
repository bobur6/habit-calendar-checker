package handlers

import (
	"github.com/gin-gonic/gin"
	"go-rest-project/internal/models"
	"go-rest-project/internal/services"
	"strconv"
	"time"
)

type HabitCheckHandler struct {
	habitCheckService services.HabitCheckService
}

func NewHabitCheckHandler(habitCheckService services.HabitCheckService) *HabitCheckHandler {
	return &HabitCheckHandler{habitCheckService: habitCheckService}
}

func (h *HabitCheckHandler) CreateHabitCheck(c *gin.Context) {
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
	var check models.HabitCheck
	if err := c.BindJSON(&check); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	// Если в теле нет habit_id, но есть в query, подставим
	if check.HabitID == 0 {
		habitIDStr := c.Query("habit_id")
		if habitIDStr != "" {
			habitID, err := strconv.ParseUint(habitIDStr, 10, 32)
			if err == nil && habitID > 0 {
				check.HabitID = uint(habitID)
			}
		}
	}

	if err := h.habitCheckService.CreateHabitCheck(&check); err != nil {
		// Можно добавить кастомные коды ошибок, если потребуется
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, check)
}

func (h *HabitCheckHandler) GetHabitCheck(c *gin.Context) {
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
		c.JSON(400, gin.H{"error": "Invalid check ID"})
		return
	}

	check, err := h.habitCheckService.GetHabitCheckByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Check not found"})
		return
	}

	c.JSON(200, check)
}

func (h *HabitCheckHandler) GetHabitChecks(c *gin.Context) {
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
	habitIDStr := c.Query("habit_id")
	if habitIDStr == "" {
		c.JSON(400, gin.H{"error": "Missing habit_id parameter"})
		return
	}
	habitID, err := strconv.ParseUint(habitIDStr, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid habit ID"})
		return
	}

	checks, err := h.habitCheckService.GetHabitChecksByHabitID(uint(habitID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, checks)
}

func (h *HabitCheckHandler) GetHabitCheckByDate(c *gin.Context) {
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
	habitIDStr := c.Query("habit_id")
	dateStr := c.Query("date")

	habitID, err := strconv.ParseUint(habitIDStr, 10, 32)
	if err != nil || habitID == 0 {
		c.JSON(400, gin.H{"error": "Invalid habit ID"})
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	check, err := h.habitCheckService.GetHabitCheckByHabitIDAndDate(uint(habitID), date)
	if err != nil {
		c.JSON(404, gin.H{"error": "Check not found"})
		return
	}

	c.JSON(200, check)
}

func (h *HabitCheckHandler) UpdateHabitCheck(c *gin.Context) {
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
		c.JSON(400, gin.H{"error": "Invalid check ID"})
		return
	}

	var check models.HabitCheck
	if err := c.BindJSON(&check); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Проверка совпадения ID
	if check.ID != 0 && check.ID != uint(id) {
		c.JSON(400, gin.H{"error": "ID in body does not match ID in URL"})
		return
	}
	check.ID = uint(id)
	if err := h.habitCheckService.UpdateHabitCheck(&check); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, check)
}

func (h *HabitCheckHandler) DeleteHabitCheck(c *gin.Context) {
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
		c.JSON(400, gin.H{"error": "Invalid check ID"})
		return
	}

	if err := h.habitCheckService.DeleteHabitCheck(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Habit check deleted successfully"})

}
