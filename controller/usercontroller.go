package controller

import (
	"lms/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	us  services.UserServices
	ctx *gin.Context
}

func (uc UserController) VerifikasiRegister(ctx *gin.Context) {
	data, err := uc.us.Register(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "error",
			"result":  data,
		})
	}

	ctx.JSON(200, gin.H{
		"message": "succesfully",
		"result":  data,
	})
}

func (uc UserController) VerifikasiLogin(ctx *gin.Context) {
	data, err := uc.us.Login(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "error",
			"result":  data,
		})
	}

	ctx.JSON(200, gin.H{
		"message": "succesfully",
		"result":  data,
	})

}

func NewUserController(us services.UserServices, ctx *gin.Context) UserController {
	return UserController{
		us:  us,
		ctx: ctx,
	}
}
