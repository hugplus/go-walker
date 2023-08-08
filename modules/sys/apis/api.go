package apis

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/common/resp"
	"github.com/hugplus/go-walker/modules/sys/service"
	"github.com/hugplus/go-walker/modules/sys/service/dto"
)

type SysApi struct {
	//api.BaseApi
}

// Ping Ping接口
// @Summary Ping接口
// @Description Ping接口
// @Tags Default
// @Success 200 {object} response.Response{data=string}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ping [get]
func (e SysApi) Ping(c *gin.Context) {
	req := dto.SysDto{}
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
		return
	}

	if err := service.Sys.PingS(&req); err != nil {
		fmt.Println(err)
		resp.Fail(c, 500, err.Error())
		return
	}
	fmt.Println("aaaa")
	resp.Ok(c, time.Now())
	//c.JSON(200, "OK")
	fmt.Println("bbbb")
}
