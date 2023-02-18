package main

import (
	"go/docs"
	"go/router"
	"go/utils"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := router.Router()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 不写127.0.0.1 win每次执行会被防火墙拦截
	r.Run("127.0.0.1:8080")
}
