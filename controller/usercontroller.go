package controller

import (
	"fmt"
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
		ctx.Abort()
		return
	}

	ctx.HTML(200,"login.html",nil)
}

func (uc UserController) VerifikasiLogin(ctx *gin.Context) {
	data, err := uc.us.Login(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "error",
			"result":  data,
		})
	}

	fmt.Println("di controller", err)

	ctx.HTML(200,"home.html",gin.H{
		"name":data.FirstName,
	})

}

func NewUserController(us services.UserServices, ctx *gin.Context) UserController {
	return UserController{
		us:  us,
		ctx: ctx,
	}
}
