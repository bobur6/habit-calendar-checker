package models

import (
	"time"
)

type User struct {
	ID          uint       `gorm:"primaryKey"`
	Username    string    `json:"username" gorm:"unique;not null"`
	Email       string    `json:"email" gorm:"unique;not null"`
	Password    string    `json:"-" gorm:"not null"`
	HabitLists  []HabitList `json:"habit_lists" gorm:"foreignKey:UserID;references:ID"`
	LastLoginAt *time.Time `json:"last_login_at"`
	Role        string    `json:"role" gorm:"default:'user'"`
}
