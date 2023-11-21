package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"xapi/api"
	_ "xapi/docs"
	"xapi/middleware"
	"xapi/service"
)

func InitRouter() {
	router := gin.Default()
	router.Use(middleware.CrosMiddleWare())
	// 使用 Swagger 中间件
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api.SetupUserRouter(router)

	err := router.Run(fmt.Sprintf(":%s", service.GetConfig().Server.Port))
	if err != nil {
		log.Println("服务器端口绑定失败", err)
	}
}
