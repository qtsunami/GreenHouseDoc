package jwt

import (
	"github.com/gin-gonic/gin"
	"GhdApi/pkg/exception"
	"GhdApi/pkg/util"
)

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
			} else if time.Now().Unix() > claims.expiresAt {
				
			}
		}

	}

}

