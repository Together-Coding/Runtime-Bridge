package runtimes

import (
	"github.com/together-coding/runtime-bridge/db"
	"time"
)

type RuntimeImage struct {
	ID           int64     `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"not null"`
	LanguageName string    `json:"language_name" gorm:"not null"` // ForeignKey
	Taskdef      string    `json:"-" gorm:"not null"`             // AWS ECS Task Definition name
	Revision     string    `json:"-" gorm:"not null"`             // Revision number of Task Definition
	Available    bool      `json:"available" gorm:"not null;default:0"`
	CreatedAt    time.Time `json:"-" gorm:"not null"`
}

type SupportedLanguage struct {
	Name  string `json:"name" gorm:"primaryKey"`
	Order int8   `json:"order" gorm:"not null;default:0"`

	RuntimeImages []RuntimeImage `json:"-" gorm:"constraint:OnDelete:CASCADE;foreignKey:LanguageName;references:Name"`
}

func GetRuntimeImage(id int64) RuntimeImage {
	ri := RuntimeImage{ID: id}
	db.DB.Where(&ri).Find(&ri)
	return ri
}
