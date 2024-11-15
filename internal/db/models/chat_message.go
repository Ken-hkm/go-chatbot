// chat_message.go
package models

import (
	"gorm.io/gorm"
	"time"
)

type ChatMessage struct {
	gorm.Model
	UserID    uint      `gorm:"not null" json:"user_id"`
	Message   string    `gorm:"type:text;not null" json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
