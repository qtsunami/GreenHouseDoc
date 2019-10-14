package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"GhdApi/pkg/exception"
	"GhdApi/pkg/util"
)

// JWT 中间件验证
func JWT() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var code int
		var data interface{}

		code = exception.SUCCESS
		token := ctx.Query("token")

		if token == "" {
			code = exception.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = exception.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = exception.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != exception.SUCCESS {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg": exception.GetMsg(code),
				"data": data,
			})
			ctx.Abort()
			return
		}

		ctx.Next()

	}

}

