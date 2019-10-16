package apis

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/astaxie/beego/validation"
	"GhdApi/pkg/exception"
	"GhdApi/pkg/util"
    "GhdApi/models"
    "GhdApi/pkg/logging"
)

type auth struct {
	Username string `valid:"Required", MaxSize(50)`
	Password string `valid:"Required", MaxSize(50)`
}

// GetAuth godoc
// @Summary List accounts
// @Description get accounts
// @Accept  json
// @Produce  json
// @Param username query string false "username"
// @Param password query string false "password"
// @Header 200 {string} Token "qwerty"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth [get]
func GetAuth(ctx *gin.Context) {
	username := ctx.Query("username")
    password := ctx.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	
	data := make(map[string]interface{})
    code := exception.INVALID_PARAMS
    if ok {
        isExist := models.CheckAuth(username, password)
        if isExist {
            token, err := util.GenerateToken(username, password)
            if err != nil {
                code = exception.ERROR_AUTH_TOKEN
            } else {
                data["token"] = token
                
                code = exception.SUCCESS
            }

        } else {
            code = exception.ERROR_AUTH
        }
    } else {
        for _, err := range valid.Errors {
            // log.Println(err.Key, err.Message)
            logging.Info(err.Key, err.Message)
        }
    }

    ctx.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : exception.GetMsg(code),
        "data" : data,
    })
}