package api

import (
	"bitbucket.org/udevs/api_gateway/config"
	"bitbucket.org/udevs/api_gateway/pkg/logger"
	"bitbucket.org/udevs/api_gateway/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// @Summary 登录
	// @Description 登录
	// @Produce json
	// @Param body body controllers.LoginParams true "body参数"
	// @Success 200 {string} string "ok" "返回用户信息"
	// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
	// @Failure 401 {string} string "err_code：10001 登录失败"
	// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
	// @Router /user/person/login [post]
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "bitbucket.org/udevs/api_gateway/api/docs"
	v1 "bitbucket.org/udevs/api_gateway/api/handlers/v1"
)

type RouterOptions struct {
	Log      logger.Logger
	Cfg      config.Config
	Services services.ServiceManager
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	router.Use(cors.New(config))
	// router.Use(MaxAllowed(100))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:      opt.Log,
		Cfg:      opt.Cfg,
		Services: opt.Services,
	})

	router.GET("/config", handlerV1.GetConfig)

	apiV1 := router.Group("/v1")
	apiV1.GET("/ping", handlerV1.Ping)

	// profession
	apiV1.POST("/profession", handlerV1.CreateProfession)
	apiV1.GET("/profession", handlerV1.GetAllProfession)
	apiV1.GET("/profession/:id", handlerV1.GetByIdProfession)
	apiV1.PUT("/profession", handlerV1.UpdateProfession)
	apiV1.DELETE("/profession", handlerV1.DeleteProfession)
	// attribute
	apiV1.POST("/attribute", handlerV1.CreateAttribute)
	apiV1.GET("/attribute", handlerV1.GetAllAttribute)
	apiV1.GET("/attribute/:id", handlerV1.GetById)
	apiV1.PUT("/attribute", handlerV1.UpdateAttribute)
	apiV1.DELETE("/attribute", handlerV1.DeleteAttribute)
	// position
	apiV1.POST("/position", handlerV1.CreatePosition)
	apiV1.GET("/position", handlerV1.GetAllPosition)
	apiV1.GET("/position/:id", handlerV1.GetByIdPosition)
	apiV1.PUT("/position", handlerV1.UpdatePosition)
	apiV1.DELETE("/position", handlerV1.DeletePosition)
	// company
	apiV1.POST("/company", handlerV1.Create)
	apiV1.GET("/company", handlerV1.GetAll)
	apiV1.GET("/company/:id", handlerV1.GetByIdCompany)
	apiV1.PUT("/company", handlerV1.Update)
	apiV1.DELETE("/company", handlerV1.Delete)

	// swagger
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

func MaxAllowed(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()

	}
}
