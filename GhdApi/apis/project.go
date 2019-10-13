package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


// List 获取项目列表
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