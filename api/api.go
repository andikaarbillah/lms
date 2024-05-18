package api

import (
	"lms/controller"
	"lms/repository"
	"lms/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var(
	ctx *gin.Context
)

func Api(router *gin.Engine, db *gorm.DB){
	userRepo := repository.NewUserRepository(db)
	userServices := services.NewUserService(userRepo)
	userController := controller.NewUserController(userServices, ctx)

	router.GET("/",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message":"welcome in web server",
		})
	})

	v1 := router.Group("/api/v1")

	{
		v1.POST("/user/register",userController.VerifikasiRegister)
		v1.POST("/user",userController.VerifikasiLogin)
	}
}