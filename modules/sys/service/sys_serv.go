package service

import (
	"time"

	"github.com/hugplus/go-walker/common/base"
	"github.com/hugplus/go-walker/core"
	"github.com/hugplus/go-walker/modules/sys/models"
	"go.uber.org/zap"
)

type SysServ struct {
	base.BaseService
}

func (s *SysServ) Ping(reqId string, d *models.Sys) error {

	cstr, err := core.Cache.Get("test")
	if err != nil || cstr == "" {
		if err := core.Cache.Set("test", d, time.Hour); err != nil {
			core.Log.Error("REQID:"+reqId, zap.Error(err))
			return err
		}
	}

	if err := core.DB().Create(&d).Error; err != nil {
		core.Log.Error("REQID:"+reqId, zap.Error(err))
		return err
	}
	return nil
}
