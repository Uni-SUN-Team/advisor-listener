package route

import (
	"unisun/api/advisor-listener/src/controller"

	"github.com/gin-gonic/gin"
)

func Consumer(r *gin.RouterGroup) {
	r.GET("/advisors", controller.Advisors)
	r.GET("/advisors/:id", controller.AdvisorById)
	r.GET("/advisors/slug/:slug", controller.AdvisorBySlug)
}
