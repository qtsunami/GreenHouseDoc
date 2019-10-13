package routers

import (
	"GhdApi/apis"
	"GhdApi/pkg/settings"
	// "net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter Router group
func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	gin.SetMode(settings.RunMode)



	project := router.Group("/project")
	{
		project.GET("/list", apis.List)
	}

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
            "message": "testsssss",
        })
	})

	

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
