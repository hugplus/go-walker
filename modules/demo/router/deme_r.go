package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/modules/demo/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, sysNoCheckRoleRouter)
}

// 无需认证的路由示例
func sysNoCheckRoleRouter(v1 *gin.RouterGroup) {
	api := apis.DemoApi{}
	r := v1.Group("demo")
	{
		r.POST("/get", api.Get)
		r.POST("/create", api.Create)
		r.POST("/update", api.Update)
		r.POST("/page", api.QueryPage)
		r.POST("/del", api.Del)
	}
}
