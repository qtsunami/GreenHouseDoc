package routers

import (
	"GhdApi/apis"
	"GhdApi/pkg/settings"
	// "net/http"

	"github.com/gin-gonic/gin"
	"GhdApi/middleware/jwt"
)

// InitRouter Router group
func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	gin.SetMode(settings.RunMode)


	router.GET("/auth", apis.GetAuth)

	project := router.Group("/project")
	project.Use(jwt.JWT())
	{
		project.GET("/list", apis.List)
	}

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
            "message": "test",
        })
	})

	

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
