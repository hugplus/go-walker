package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/common/api"
	"github.com/hugplus/go-walker/common/api/resp"
	"github.com/hugplus/go-walker/modules/sys/models"
	"github.com/hugplus/go-walker/modules/sys/service"
	"github.com/hugplus/go-walker/modules/sys/service/dto"
)

type DemoApi struct {
	api.BaseApi
}

// Ping Ping接口
// @Summary Ping接口
// @Description Ping接口
// @Tags Default
// @Success 200 {object} response.Response{data=string}} "{"code": 200, "data": [...]}"
// @Router /api/v1/demo/ping [get]
func (e *DemoApi) Ping(c *gin.Context) {
	req := dto.SysDto{}
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Sys
	if err := service.Sys.Ping(e.GetReqId(c), &req, &data); err != nil {
		e.Error(c, err)
		return
	}

	resp.Ok(c, data)
}
