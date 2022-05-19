package src

import (
	"os"
	"unisun/api/advisor-listener/docs"
	"unisun/api/advisor-listener/src/constants"
	"unisun/api/advisor-listener/src/controller"
	"unisun/api/advisor-listener/src/route"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @termsOfService  http://swagger.io/terms/

// @contact.name    API Support
// @contact.url     http://www.swagger.io/support
// @contact.email   support@swagger.io

// @license.name  MIT License Copyright (c) 2022 Uni-SUN-Team
// @license.url   https://api.unisun.dynu.com/article-listenner/api/license

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func App() *gin.Engine {
	docs.SwaggerInfo.Title = "COURSE LISTENER API"
	docs.SwaggerInfo.Description = ""
	docs.SwaggerInfo.Version = os.Getenv(constants.VERSION)
	docs.SwaggerInfo.Host = os.Getenv(constants.HOST)
	docs.SwaggerInfo.BasePath = os.Getenv(constants.CONTEXT_PATH) + "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	g := r.Group(os.Getenv(constants.CONTEXT_PATH) + "/api")
	{
		g.GET("/healcheck", controller.HealthCheckHandler)
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		g.StaticFile("/license", "./LICENSE")
		route.Consumer(g)
	}
	return r
}
