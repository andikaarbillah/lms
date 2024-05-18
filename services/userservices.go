package services

import (
	"lms/hashing"
	"lms/model"
	"lms/primary"
	"lms/repository"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserServices interface {
	Register(ctx *gin.Context) (*model.Users, error)
	Login(ctx *gin.Context) (*model.Users, error)
}

type userServices struct {
	ur repository.UserRepository
}

func (ur *userServices) Register(ctx *gin.Context) (*model.Users, error) {
	var user model.Users

	if err := ctx.ShouldBindJSON(&user); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(user)

	if err != nil {
		return nil, err
	}

	User := model.Users{
		Id:         primary.IdRndm(8),
		Username:   user.Username,
		Email:      user.Email,
		Password:   hashing.HasingString(user.Password),
		Created_at: time.Now(),
		Updated_at: time.Now(),
		Is_deleted: false,
	}

	result, err := ur.ur.Register(User)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (ur *userServices) Login(ctx *gin.Context) (*model.Users, error) {
	var user model.Users
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return nil, err
	}
	validate := validator.New()
	err := validate.Var(user.Email,"required,email")
	if err != nil {
		return nil, err
	}

	err = validate.Var(user.Password,"required")
	if err != nil {
		return nil, err
	}

	existingUser, err := ur.ur.Login(user.Email, user.Password)

	if err != nil{
		return nil, err
	}

	return existingUser, nil
}

func NewUserService(ur repository.UserRepository) UserServices {
	return &userServices{
		ur: ur,
	}
}
