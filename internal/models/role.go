package models

import "time"

type Role struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex;size:255;not null" json:"name"`
	Level       int       `gorm:"not null" json:"level"`
	Description string    `gorm:"size:1000" json:"description"`
	CreatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"-"`
}
