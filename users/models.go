package users

import (
	"github.com/together-coding/runtime-bridge/containers"
	"time"
)

// Note that, `users` table is managed by "Authentication-Server"

type User struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Email      string    `json:"-" gorm:"not null"`
	Password   string    `json:"-" gorm:"not null"`
	Name       string    `json:"-" gorm:"not null"`
	FromSocial bool      `json:"-" gorm:"not null"`
	CreatedAt  time.Time `json:"-" gorm:"not null"`

	RuntimeAllocations []containers.RuntimeAllocation `gorm:"constraint:OnDelete:CASCADE"`
}
