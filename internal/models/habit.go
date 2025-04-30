package models

import (
	"time"
)

type Habit struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	HabitListID uint        `json:"habit_list_id" gorm:"not null"`
	Name        string      `json:"name" gorm:"not null"`
	Description string      `json:"description"`
	DefaultEmoji string    `json:"default_emoji" gorm:"default:'✅'"`
	Checks      []HabitCheck `json:"checks" gorm:"foreignKey:HabitID;references:ID"`
	IsArchived  bool        `json:"is_archived" gorm:"default:false"`
}