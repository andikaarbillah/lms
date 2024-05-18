package repository

import (
	"errors"
	"lms/hashing"
	"lms/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user model.Users) (*model.Users, error)
	Login(email string, password string) (*model.Users, error)
}

type userRepository struct {
	db *gorm.DB
}

func (ur *userRepository) Register(user model.Users) (*model.Users, error) {
	result := ur.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (ur *userRepository) Login(email string, password string) (*model.Users, error) {
	var user model.Users
	result := ur.db.Where("email=?", email).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	if !hashing.EncriptPassword(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	return &user, nil

}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
