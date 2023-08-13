package service

import (
	"github.com/hugplus/go-walker/common/consts"
	"github.com/hugplus/go-walker/common/errs"
	"github.com/hugplus/go-walker/common/errs/codes"
	"github.com/hugplus/go-walker/core"
	"github.com/hugplus/go-walker/modules/demo/models"
	"go.uber.org/zap"
)

type DemoService struct {
}

func (*DemoService) Ping(reqId string, d2 *models.Demo) errs.IError {
	if err := core.Db(consts.DB_DEMO).Create(&d2).Error; err != nil {
		berr := errs.Err(codes.FAILURE, reqId, err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	return nil
}
