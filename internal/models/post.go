package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        int64          `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"size:255;not null" json:"title"`
	Content   string         `gorm:"size:10000;not null" json:"content"`
	UserID    int64          `gorm:"not null" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user"`
	Tags      []Tag          `gorm:"many2many:post_tags;" json:"tags"`
	CreatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Comments  []Comment      `gorm:"foreignKey:PostID" json:"comments,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Tag struct {
	ID   int64  `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique;not null" json:"name"`
}
