package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"

	"casorder/db/models"
	"casorder/utils/types"
)

// Authorized blocks unauthorized requestrs
func Authorized(c *gin.Context) {
	_, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(401)
		return
	}
}

func validateToken(token string) (types.JSON, error) {
	// TODO handle validate manageiq session
	return nil, nil
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			authorization := c.Request.Header.Get("Authorization")
			if authorization == "" {
				c.Next()
				return
			}
			sp := strings.Split(authorization, "Bearer ")
			// invalid token
			if len(sp) < 1 {
				c.Next()
				return
			}
			token = sp[1]
		}

		tokenData, err := validateToken(token)
		if err != nil {
			c.Next()
			return
		}

		var user models.User
		user.Read(tokenData["user"].(types.JSON))

		c.Set("user", user)
		c.Set("token_expire", tokenData["exp"])
		c.Next()
	}
}
