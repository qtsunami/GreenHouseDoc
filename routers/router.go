package routers

import (
	"GhdApi/pkg/settings"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter Router group
func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	gin.SetMode(settings.RunMode)

	router.Delims("{<", ">}")

	router.LoadHTMLGlob("pages/**/*")
	router.Static("/static", "./static")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home/index.html", gin.H{
			"title": "Home Page",
		})
	})

	router.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login/login.html", gin.H{
			"title": "Login Page",
		})
	})

	router.GET("/register", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login/register.html", gin.H{
			"title": "Login Page",
		})
	})

	router.Post("/register")

	router.GET("/v1", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home/index1.html", gin.H{
			"title": "Home Page",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
