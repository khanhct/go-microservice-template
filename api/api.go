package api

import (
	"github.com/gin-gonic/gin"

	apiv1 "casorder/api/v1"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/casorder/api")
	{
		apiv1.ApplyRoutes(api)
	}
}
