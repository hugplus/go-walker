package apis

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/common/api"
	"github.com/hugplus/go-walker/common/resp"
)

type SysApi struct {
	api.BaseApi
}

// Ping Ping接口
// @Summary Ping接口
// @Description Ping接口
// @Tags Default
// @Success 200 {object} response.Response{data=string}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ping [get]
func (e *SysApi) Ping(c *gin.Context) {
	//e.GetReqId(c)
	resp.Ok(c, time.Now())
}
