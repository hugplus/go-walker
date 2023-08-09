package service

import (
	"time"

	"github.com/hugplus/go-walker/common/service"
	"github.com/hugplus/go-walker/core"
	"github.com/hugplus/go-walker/modules/sys/models"
	"github.com/hugplus/go-walker/modules/sys/service/dto"
	"go.uber.org/zap"
)

type SysServ struct {
	service.BaseService
}

func (s *SysServ) Ping(reqId string, req *dto.SysDto, d *models.Sys) error {

	d2 := models.Sys{
		Name: req.Name,
	}

	cstr, err := core.Cache.Get("test")
	if err != nil || cstr == "" {
		if err := core.Cache.Set("test", d2, time.Hour); err != nil {
			core.Log.Error("REQID:"+reqId, zap.Error(err))
			return err
		}
	}

	if err := core.Db("master").Create(&d2).Error; err != nil {
		core.Log.Error("REQID:"+reqId, zap.Error(err))
		return err
	}
	*d = d2
	return nil
}
