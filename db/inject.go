package db

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// Inject injects database to gin context
func Inject(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
