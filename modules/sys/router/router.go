package router

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/core"
	//common "go-admin/common/middleware"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	//routerCheckRole   = make([]func(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware), 0)
)

// InitRouter 路由初始化，不要怀疑，这里用到了
func InitRouter() {
	var r *gin.Engine
	h := core.GetEngine()
	if h == nil {
		log.Fatal("not found engine...")
		os.Exit(-1)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}

	// the jwt middleware
	// authMiddleware, err := common.AuthInit()
	// if err != nil {
	// 	log.Fatalf("JWT Init Error, %s", err.Error())
	// }

	// // 注册系统路由
	// InitSysRouter(r, authMiddleware)

	// // 注册业务路由
	// // TODO: 这里可存放业务路由，里边并无实际路由只有演示代码
	// InitExamplesRouter(r, authMiddleware)
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