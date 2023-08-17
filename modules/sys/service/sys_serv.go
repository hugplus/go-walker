package service

import (
	"fmt"
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

func (s *SysServ) Init() errs.IError {
	core.DB().AutoMigrate(
		&models.Sys{},
	)

	cstr, err := core.Cache.Get("test")
	fmt.Printf("Init %s ,%v \n", cstr, err)
	//if err != nil || cstr == "" {
	d := models.Sys{
		Name: "goods",
	}
	d.UpdatedAt = time.Now().Unix()
	d.CreatedAt = d.UpdatedAt
	d.Status = 3
	d.UpdateBy = 3
	d.CreateBy = d.UpdateBy

	if err := core.DB().Create(&d).Error; err != nil {
		berr := errs.Err(codes.FAILURE, "reqId", err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	if err := core.Cache.Set("test", d, time.Hour); err != nil {
		berr := errs.Err(codes.FAILURE, "reqId", err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	return nil
}
