package base

import (
	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/common/base/resp"
	"github.com/hugplus/go-walker/common/consts"
	"github.com/hugplus/go-walker/common/utils"
)

type BaseApi struct {
}

func (e *BaseApi) GetReqId(c *gin.Context) string {
	return utils.GetReqId(c)
}

func (e *BaseApi) GetUserId(c *gin.Context) int {
	return c.GetInt(consts.USER_ID)
}

func (e *BaseApi) GetTenantId(c *gin.Context) int {
	return c.GetInt(consts.TENANT_ID)
}

func (e *BaseApi) Error(c *gin.Context, err error) {
	//e.LogError(c, err)
	resp.Fail(c, resp.ERROR_CODE, err.Error())
}

//封装后代码路径指定到这里所以去掉
// func (e *BaseApi) LogError(c *gin.Context, err error) {
// 	core.Log.Error(fmt.Sprintf("REQID:%s", e.GetReqId(c)), zap.Error(err))
// }

// func (e *BaseApi) LogInfo(c *gin.Context, key string, val any) {
// 	ccore.Log.Info("REQID"+e.GetReqId(c), zap.Reflect("data", data))
// }
