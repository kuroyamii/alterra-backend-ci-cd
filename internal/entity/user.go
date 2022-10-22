package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primarykey,autoincrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Email     string
	Password  string
}

type Users []User
