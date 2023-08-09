package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/common/api"
	"github.com/hugplus/go-walker/common/api/resp"
	"github.com/hugplus/go-walker/common/utils"
	"github.com/hugplus/go-walker/core"
	"go.uber.org/zap"
)

type SysApi struct {
	api.BaseApi
}

// Ping Ping接口
// @Summary Ping接口
// @Description Ping接口
// @Tags Default
// @Success 200 {object} response.Response{data=utils.Server}} "{"code": 200, "data": [...]}"
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
	resp.Ok(c, server)
}
