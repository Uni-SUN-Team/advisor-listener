package controller

import "github.com/gin-gonic/gin"

type ControllerPort interface {
	Advisors(c *gin.Context)
	AdvisorById(c *gin.Context)
	AdvisorBySlug(c *gin.Context)
}

type ControllerHealPort interface {
	HealthCheckHandler(c *gin.Context)
}
