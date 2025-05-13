package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        int64          `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"uniqueIndex;size:255;not null" json:"username"`
	Email     string         `gorm:"uniqueIndex;size:255;not null" json:"email"`
	Password  []byte         `gorm:"size:255;not null" json:"-"`
	CreatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	IsActive  bool           `gorm:"not null;default:false" json:"is_active"`
	RoleID    int64          `gorm:"not null" json:"role_id"`
	Role      Role           `gorm:"foreignKey:RoleID" json:"role"`
	Posts     []Post         `gorm:"foreignKey:UserID" json:"posts,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *User) SetPassword(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
}

func (u *User) CheckPassword(text string) bool {
	err := bcrypt.CompareHashAndPassword(u.Password, []byte(text))
	return err == nil
}
