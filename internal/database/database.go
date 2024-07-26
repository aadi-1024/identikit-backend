package database

import (
	"github.com/aadi-1024/identikit-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	conn *gorm.DB
}

func InitDb(dsn string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	return &Database{db}, nil
}
