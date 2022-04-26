package runtimes

import (
	"github.com/together-coding/runtime-bridge/containers"
	"time"
)

type RuntimeImage struct {
	ID           int64     `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"not null"`
	LanguageName string    `json:"language_name" gorm:"not null"` // ForeignKey
	Taskdef      string    `json:"taskdef" gorm:"not null"`       // AWS ECS Task Definition name
	Revision     string    `json:"revision" gorm:"not null"`      // Revision number of Task Definition
	Available    bool      `json:"available" gorm:"not null;default:0"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null"`

	RuntimeAllocations []containers.RuntimeAllocation `gorm:"constraint:OnDelete:CASCADE"`
}

type SupportedLanguage struct {
	Name  string `json:"name" gorm:"primaryKey"`
	Order int8   `json:"order" gorm:"not null;default:0"`

	RuntimeImages []RuntimeImage `json:"images" gorm:"constraint:OnDelete:CASCADE;foreignKey:LanguageName;references:Name"`
}
