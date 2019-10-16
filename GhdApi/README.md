# GreenHouseDoc 文档管理平台
> 主要用于日常项目wiki文档管理


### Swagger api documentation

**参考文档**
- https://github.com/swaggo/swag

**安装 Swagger**
```bash
$ go get -u github.com/swaggo/swag/cmd/swag
```

**验证Swagger是否安装成功**
```bash
$ swag -v
```

**安装 gin-swagger**
```
$ go get -u github.com/swaggo/gin-swagger
$ go get -u github.com/swaggo/gin-swagger/swaggerFiles
```

**编写API注释**
example:
```
// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} model.Account
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts/{id} [get]
```

**生成文档**
```
$ swag init
```
完成后，会在根目录下生成docs目录：

```
docs/
├── docs.go
├── swagger.json
└── swagger.yaml
```

**验证访问**
- 访问地址： http://localhost:8000/swagger/index.html