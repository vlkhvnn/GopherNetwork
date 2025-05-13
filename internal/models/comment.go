package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int64          `gorm:"primaryKey" json:"id"`
	Content   string         `gorm:"size:1000;not null" json:"content"`
	UserID    int64          `gorm:"not null" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user"`
	PostID    int64          `gorm:"not null" json:"post_id"`
	Post      Post           `gorm:"foreignKey:PostID" json:"post"`
	CreatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
