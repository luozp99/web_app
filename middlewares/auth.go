package auth

import (
	"github.com/gin-gonic/gin"
	"strings"
	"web_app/controllers"
	"web_app/pkg/jwt"
)

func JwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controllers.ResponseError(c, controllers.CodeInvalidationLogin)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controllers.ResponseError(c, controllers.CodeInvalidationTokenEmpty)
			c.Abort()
			return
		}
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controllers.ResponseError(c, controllers.CodeInvalidationToken)
			c.Abort()
			return
		}
		c.Set(controllers.UserIdKey, mc.UserID)
		c.Set(controllers.UserName, mc.Username)
		c.Next()
	}
}
