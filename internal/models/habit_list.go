package models

import (
	"time"
)

type HabitList struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID     uint   `json:"user_id" gorm:"not null"`
	Name       string `json:"name" gorm:"not null"`
	Habits     []Habit `json:"habits" gorm:"foreignKey:HabitListID;references:ID"`
	IsArchived bool   `json:"is_archived" gorm:"default:false"`
}