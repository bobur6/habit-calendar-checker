package models

import (
	"time"
)

type HabitCheck struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	HabitID    uint      `json:"habit_id" gorm:"not null;uniqueIndex:idx_habit_date"`
	Date       time.Time `json:"date" gorm:"not null;uniqueIndex:idx_habit_date"`
	Emoji      string    `json:"emoji" gorm:"default:'✅'"`
	Note       *string    `json:"note"`
	IsCompleted bool     `json:"is_completed" gorm:"default:true"`
}