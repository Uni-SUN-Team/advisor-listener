package route

import (
	"unisun/api/advisor-listener/src/controller"

	"github.com/gin-gonic/gin"
)

type ConsumerRouteAdapter struct {
	Controller controller.ControllerPort
}

func NewConsumerRouteAdapter(controllerValue controller.ControllerPort) *ConsumerRouteAdapter {
	return &ConsumerRouteAdapter{
		Controller: controllerValue,
	}
}

func (srv *ConsumerRouteAdapter) Consumer(r *gin.RouterGroup) {
	r.GET("/advisors", srv.Controller.Advisors)
	r.GET("/advisors/:id", srv.Controller.AdvisorById)
	r.GET("/advisors/slug/:slug", srv.Controller.AdvisorBySlug)
}
