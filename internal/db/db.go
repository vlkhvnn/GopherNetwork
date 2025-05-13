package db

import (
	"GopherNetwork/internal/models"
	"context"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(addr string, maxOpenConns int, maxIdleConns int, maxIdleTime string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(addr), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(models.User{}, models.Post{}, models.Comment{}, models.Follower{}, models.Role{}, models.UserInvitation{})

	if err != nil {
		return nil, err
	}

	if err := seedRoles(db); err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func seedRoles(db *gorm.DB) error {
	roles := []models.Role{
		{
			Name:        "user",
			Level:       1,
			Description: "A user can create posts and comments",
		},
		{
			Name:        "moderator",
			Level:       2,
			Description: "A moderator can update other users posts and comments",
		},
		{
			Name:        "admin",
			Level:       3,
			Description: "An admin can update and delete other users posts and comments",
		},
	}

	for _, role := range roles {
		var existing models.Role
		err := db.Where("name = ?", role.Name).First(&existing).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// Create new role
				if err := db.Create(&role).Error; err != nil {
					return err
				}
			} else {
				// Other DB error
				return err
			}
		} else {
			// Update existing role
			existing.Level = role.Level
			existing.Description = role.Description
			if err := db.Save(&existing).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
