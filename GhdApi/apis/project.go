package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


// List godoc
// @Summary List project
// @Description get project list
// @Accept  json
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Router /project/list [get]
func List(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
        "code" : 0,
        "msg" : "success",
        "data" : 100010,
    })
}

// Item 获取单个项目信息
func Item(ctx *gin.Context) {

}