package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/common/base"
	"github.com/hugplus/go-walker/common/base/resp"
	"github.com/hugplus/go-walker/common/utils"
	"github.com/hugplus/go-walker/core"
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
// @Success 200 {object} resp.Resp{data=utils.Server} "{"code": 200, "data": [...]}"
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

// init init接口
// @Summary init接口
// @Description init接口
// @Tags Default
// @Success 200 {object} resp.Resp{data=string} "{"code": 200, "data": [...]}"
// @Router /api/v1/init [get]
func (e *SysApi) Init(c *gin.Context) {
	// core.DB().AutoMigrate(
	// // models.CasbinRule{},
	// // models.DictData{},
	// // models.DictType{},
	// // models.SysRole{},
	// // models.SysApi{},
	// // models.SysConfig{},
	// // models.SysDept{},
	// // models.SysColumns{},
	// // models.SysJob{},
	// // models.SysLoginLog{},
	// // models.SysMenu{},
	// // models.SysOperaLog{},
	// // models.SysPost{},
	// // models.SysRoleDept{},
	// // models.SysTables{},
	// // models.SysUser{},
	// // models.SysPost{},
	// // models.SysLoginLog{},
	// )
	var req dto.SysDto
	if err := c.ShouldBind(&req); err != nil {

		e.Error(c, err)
		return
	}
	var data models.Sys
	data.Name = req.Name
	if err := service.Sys.Ping(e.GetReqId(c), &data); err != nil {
		resp.Fail(c, 500, "错误了")
		return
	}
	resp.Ok(c, data)
}
