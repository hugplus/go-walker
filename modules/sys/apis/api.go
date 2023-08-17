package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/common/utils"
	"github.com/hugplus/go-walker/core"
	"github.com/hugplus/go-walker/core/base"
	"github.com/hugplus/go-walker/modules/sys/models"
	"github.com/hugplus/go-walker/modules/sys/service"
	"github.com/hugplus/go-walker/modules/sys/service/dto"
	"go.uber.org/zap"
)

type SysApi struct {
	base.BaseApi
}

// Ping Ping接口
// @Summary Ping接口
// @Description Ping接口
// @Tags Default
// @Success 200 {object} base.Resp{data=utils.Server} "{"code": 200, "data": [...]}"
// @Router /api/v1/ping [get]
func (e *SysApi) Ping(c *gin.Context) {
	cpu, err := utils.InitCPU()
	if err != nil {
		core.Log.Error("Get CPU ERR", zap.Error(err))
	}
	ram, err := utils.InitRAM()
	if err != nil {
		core.Log.Error("Get RAM ERR", zap.Error(err))
	}
	disk, err := utils.InitDisk()
	if err != nil {
		core.Log.Error("Get DISK ERR", zap.Error(err))
	}
	server := utils.Server{
		Os:   utils.InitOS(),
		Cpu:  cpu,
		Ram:  ram,
		Disk: disk,
	}
	e.Ok(c, server)
}

// init init接口
// @Summary init接口
// @Description init接口
// @Tags Default
// @Success 200 {object} base.Resp{data=string} "{"code": 200, "data": [...]}"
// @Router /api/v1/init [get]
func (e *SysApi) Init(c *gin.Context) {
	var req dto.SysDto
	if err := c.ShouldBind(&req); err != nil {
		return
	}
	var data models.Sys
	data.Name = req.Name
	if err := service.Sys.Init(); err != nil {
		e.Err(c, err)
		return
	}
	e.Ok(c)
}
