package database

import (
	"context"

	"github.com/aadi-1024/identikit-backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (d *Database) RegisterUser(ctx context.Context, user models.User) (int, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), -1)
	if err != nil {
		return 0, err
	}
	user.Password = string(pass)

	err = d.conn.WithContext(ctx).Create(&user).Error
	return user.Id, err
}

func (d *Database) LoginUser(ctx context.Context, email, password string) (models.User, error) {
	original := models.User{}

	err := d.conn.WithContext(ctx).Find(&original).Where("email = ?", email).Error
	if err != nil {
		return original, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(original.Password), []byte(password))
	original.Password = ""

	return original, err
}
