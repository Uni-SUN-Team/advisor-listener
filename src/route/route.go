package route

import "github.com/gin-gonic/gin"

type RoutePort interface {
	Consumer(r *gin.RouterGroup)
}
