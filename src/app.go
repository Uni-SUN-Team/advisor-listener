package src

import (
	"unisun/api/advisor-listener/docs"
	"unisun/api/advisor-listener/src/controller"
	"unisun/api/advisor-listener/src/route"
	"unisun/api/advisor-listener/src/service"
	"unisun/api/advisor-listener/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
	docs.SwaggerInfo.Version = viper.GetString("app.version")
	docs.SwaggerInfo.Host = viper.GetString("app.host")
	docs.SwaggerInfo.BasePath = viper.GetString("app.context_path") + viper.GetString("app.root_path")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	controllerHeal := controller.NewHealCheckControllerAdapter()

	utilsAdap := utils.NewHTTPRequestAdapter()
	serviceAdap := service.NewConsumerServiceAdapter(utilsAdap)
	controllerAdap := controller.NewConsumerControllerAdapter(serviceAdap)
	routeAdap := route.NewConsumerRouteAdapter(controllerAdap)

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	g := r.Group(viper.GetString("app.context_path") + viper.GetString("app.root_path") + "/v1")
	{
		g.GET("/healcheck", controllerHeal.HealthCheckHandler)
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		g.StaticFile("/license", "./LICENSE")
		routeAdap.Consumer(g)
	}
	return r
}
