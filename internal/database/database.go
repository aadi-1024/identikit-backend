package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	conn *gorm.DB
}

func InitDb(dsn string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn))
	return &Database{db}, err
}
