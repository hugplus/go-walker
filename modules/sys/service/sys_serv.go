package service

import (
	"time"

	"github.com/hugplus/go-walker/common/codes"
	"github.com/hugplus/go-walker/core"
	"github.com/hugplus/go-walker/core/base"
	"github.com/hugplus/go-walker/core/errs"
	"github.com/hugplus/go-walker/modules/sys/models"
	"go.uber.org/zap"
)

type SysServ struct {
	base.BaseService
}

func (s *SysServ) Ping(reqId string, d *models.Sys) errs.IError {

	cstr, err := core.Cache.Get("test")
	if err != nil || cstr == "" {
		if err := core.Cache.Set("test", d, time.Hour); err != nil {
			berr := errs.Err(codes.FAILURE, reqId, err)
			core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
			return berr
		}
	}

	if err := core.DB().Create(&d).Error; err != nil {
		berr := errs.Err(codes.FAILURE, reqId, err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	return nil
}
