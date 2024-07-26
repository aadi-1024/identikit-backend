package database

import (
	"context"
	"time"

	"github.com/aadi-1024/identikit-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	conn    *gorm.DB
	timeout time.Duration
}

func InitDb(dsn string, timeout time.Duration) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&models.User{}, &models.Snippet{}); err != nil {
		return nil, err
	}

	return &Database{db, timeout}, nil
}

func (d *Database) context(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, d.timeout)
}
