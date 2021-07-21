package health

import (
	"github.com/gin-gonic/gin"

	"casorder/db/models"
)

// check Health of Service
func check(c *gin.Context) {

	h := models.Health{
		StatusCode: 200,
		Message:    "I'm OK",
	}
	c.JSON(200, h.Serialize())
}
