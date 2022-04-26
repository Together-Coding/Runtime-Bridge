package containers

import (
	"github.com/together-coding/runtime-bridge/runtimes"
	"github.com/together-coding/runtime-bridge/users"
	"time"
)

type RuntimeAllocation struct {
	ID     int64  `json:"id" gorm:"primaryKey"`
	UserID int64  `json:"user_id" gorm:"index;not null"` // Foreign key
	UserIp string `json:"user_ip" gorm:"not null"`

	RuntimeImageID int64     `json:"runtime_image_id" gorm:"not null"` // Foreign key
	ContLaunchedAt time.Time `json:"cont_launched_at" gorm:""`
	ContIp         string    `json:"cont_ip" gorm:""`
	ContPort       uint16    `json:"cont_port" gorm:""`
	ContUser       string    `json:"cont_user" gorm:""`
	ContAuthType   string    `json:"cont_auth_type" gorm:""`             // password, host key, etc.
	ContAuth       string    `json:"cont_auth" gorm:""`                  // authorization password. It is ok to store raw value of this?
	ContAPIKey     string    `json:"cont_api_key" gorm:"index;not null"` // Agent must verify requests whether it is from the container

	// Positive: health flag. Negative: number of health check failure.
	// 0: inactive : default or terminated status. Deleted some row data.
	// 1: launching : create container and waiting a resp from it
	// 2 : active : got resp (pong) and filled fields.
	// 3 : terminating : send kill signal
	Health int8 `json:"health" gorm:"default:0;not null"`

	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`

	// Foreign key
	User         users.User            `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	RuntimeImage runtimes.RuntimeImage `gorm:"foreignKey:RuntimeImageID;constraint:OnDelete:CASCADE;"`
}
