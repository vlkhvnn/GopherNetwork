package models

import "time"

type Follower struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	FollowerID  int64     `gorm:"not null;index:idx_follower_following,unique" json:"follower_id"`
	FollowingID int64     `gorm:"not null;index:idx_follower_following,unique" json:"following_id"`
	CreatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}
