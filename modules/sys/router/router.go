package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/core"

	"github.com/hugplus/go-walker/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	//routerCheckRole   = make([]func(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware), 0)
)

// InitRouter 路由初始化，不要怀疑，这里用到了
func InitRouter() {
	r := core.GetGinEngine()
	if core.Cfg.Server.Mode != core.ModeProd.String() {
		//初始化swagger
		fmt.Printf("Swagger %s %s start\r\n", docs.SwaggerInfo.Title, docs.SwaggerInfo.Version)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	noCheckRoleRouter(r)
}

// func InitBusinessRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {

// 	// 无需认证的路由
// 	noCheckRoleRouter(r)
// 	// 需要认证的路由
// 	checkRoleRouter(r, authMiddleware)

// 	return r
// }

// noCheckRoleRouter 无需认证的路由
func noCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v := r.Group("/api/v1")

	for _, f := range routerNoCheckRole {
		f(v)
	}
}
