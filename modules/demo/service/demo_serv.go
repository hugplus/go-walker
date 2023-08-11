package service

import (
	"github.com/hugplus/go-walker/common/base"
	"github.com/hugplus/go-walker/common/consts"
	"github.com/hugplus/go-walker/core"
	"github.com/hugplus/go-walker/modules/demo/models"
	"go.uber.org/zap"
)

type DemoService struct {
}

func (*DemoService) Ping(reqId string, d2 *models.Demo) error {

	if err := core.Db(consts.DB_DEMO).Create(&d2).Error; err != nil {
		core.Log.Error(base.FmtReqId(reqId), zap.Error(err))
		return err
	}
	return nil
}
