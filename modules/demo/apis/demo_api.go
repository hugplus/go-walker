package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/common/base"
	"github.com/hugplus/go-walker/modules/demo/models"
	"github.com/hugplus/go-walker/modules/demo/service"
	"github.com/hugplus/go-walker/modules/demo/service/dto"
)

type DemoApi struct {
	base.BaseApi
}

// Ping Ping接口
// @Summary Ping接口
// @Tags Demo
// @Accept application/json
// @Product application/json
// @Param data body dto.DemoDto true "body"
// @Success 200 {object} resp.Resp{data=models.Demo} "{"code": 200, "data": [...]}"
// @Router /api/v1/demo/ping [post]
func (e *DemoApi) Ping(c *gin.Context) {
	var req dto.DemoDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Demo
	data.Name = req.Name
	if err := service.Demo.Ping(e.GetReqId(c), &data); err != nil {
		e.Err(c, err)
		return
	}
	e.Ok(c, data)
}
