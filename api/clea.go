package api

import (
	"lms/config"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	db := config.DB()

	Api(r, db)

	r.Run(":8082")
}
