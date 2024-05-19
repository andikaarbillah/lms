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

	router.LoadHTMLGlob("template/*")
	router.Static("image","./image")
	router.Static("template","./template")

	router.GET("/",func(ctx *gin.Context) {
		ctx.HTML(200,"login.html", nil)
	})

	v1 := router.Group("/api/v1")

	{
		v1.POST("/user/register",userController.VerifikasiRegister)
		v1.POST("/user",userController.VerifikasiLogin)
	}
}