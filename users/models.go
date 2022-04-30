package users

import (
	"time"
)

// Note that, `users` table is managed by "Authentication-Server"

type User struct {
	ID         int64     `json:"id" gorm:"primaryKey"`
	Email      string    `json:"-" gorm:"not null"`
	Password   string    `json:"-" gorm:"not null"`
	Name       string    `json:"-" gorm:"not null"`
	FromSocial bool      `json:"-" gorm:"not null"`
	CreatedAt  time.Time `json:"-" gorm:"not null"`
}
