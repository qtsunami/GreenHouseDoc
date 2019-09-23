package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
}

func Register(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	nickname := ctx.PostForm("nickname")
	verifycode := ctx.PostForm("verify_code")

	ctx.JSON(http.StatusOK, gin.H{
		"email":      email,
		"password":   password,
		"nickname":   nickname,
		"verifycode": verifycode,
	})
}

func Login(ctx *gin.Context) {

}
