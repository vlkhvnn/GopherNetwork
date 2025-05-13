package models

import "time"

type UserInvitation struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	UserID    int64     `gorm:"uniqueIndex;not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	Token     string    `gorm:"uniqueIndex;size:255;not null" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}
