package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/modules/sys/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, sysNoCheckRoleRouter)
}

// 需要登录的路由
// func registerAuthCheckRoleRouter(v1 *gin.RouterGroup) {
// 	//api := apis.SSO{}
// 	r := v1.Group("/auth").Use(middleware.JWTAuthMiddleware())
// 	{
// 		r.POST("/myUserinfo", api.MyUserInfo)
// 		r.POST("/logout", api.Logout)
// 		r.POST("/changePwd", api.ChangePwd)
// 		r.POST("/bind", api.Bind)
// 		r.POST("/bindWechat", api.BindWechat)
// 		r.POST("/bindDing", api.BindDing)
// 		r.POST("/changeUserinfo", api.ChangeUserinfo)
// 	}
// }

// 无需认证的路由示例
func sysNoCheckRoleRouter(v1 *gin.RouterGroup) {
	//api := apis.SSO{}
	r := v1.Group("")
	{
		r.GET("/ping", apis.Ping)
		// r.POST("/loginWechat", api.LoginByWechat)
		// r.POST("/getDingCfg", api.GetDingCfg)
		// r.POST("/loginDing", api.LoginByDing)
		// r.POST("/register", api.Register)
		// r.POST("/captcha", api.GenerateCaptchaHandler)
		// r.POST("/sendCode", api.SendCode)
		// r.POST("/getUserinfo", api.GetUserInfo)
		// r.POST("/forgetPwd", api.ForgetPwd)

	}
}
