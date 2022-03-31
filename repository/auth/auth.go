package auth

import (
	"newtestpostgre/entities"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (a *AuthRepository) Login(email, password string) (entities.User, error) {
	user := entities.User{Email: email, Password: password}
	if err := a.db.Where("email= ? AND password= ?", email, password).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
