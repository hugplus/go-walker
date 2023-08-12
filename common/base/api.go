package base

import (
	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/common/base/resp"
	"github.com/hugplus/go-walker/common/consts"
	"github.com/hugplus/go-walker/common/errs"
	"github.com/hugplus/go-walker/common/errs/codes"
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
	resp.Fail(c, codes.FAILURE, err.Error())
}

func (e *BaseApi) WithCodeError(c *gin.Context, code int, err error) {
	resp.Fail(c, code, err.Error())
}

func (e *BaseApi) Fail(c *gin.Context, code int, msg string, data ...any) {
	resp.Fail(c, code, msg, data)
}

func (e *BaseApi) Err(c *gin.Context, err errs.IError) {
	resp.Err(c, err)
}

func (e *BaseApi) Ok(c *gin.Context, data any) {
	resp.Ok(c, data)
}

//封装后代码路径指定到这里所以去掉
// func (e *BaseApi) LogError(c *gin.Context, err error) {
// 	core.Log.Error(fmt.Sprintf("REQID:%s", e.GetReqId(c)), zap.Error(err))
// }

// func (e *BaseApi) LogInfo(c *gin.Context, key string, val any) {
// 	ccore.Log.Info("REQID"+e.GetReqId(c), zap.Reflect("data", data))
// }
