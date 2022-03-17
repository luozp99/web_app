package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/controllers"
	"web_app/logger"
	"web_app/settings"
)

func Setup(conf *settings.AppConfig) *gin.Engine {
	r := gin.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.GET("/appName", func(context *gin.Context) {
		context.String(http.StatusOK, conf.Name)
	})

	r.GET("/signup", controllers.SignUpHandler)

	r.GET("/login", controllers.LoginHandler)

	return r
}
