package service

import (
	"github.com/hugplus/go-walker/common/consts"
	"github.com/hugplus/go-walker/core"
	"github.com/hugplus/go-walker/modules/demo/models"
	"go.uber.org/zap"
)

type DemoService struct {
}

func (*DemoService) Ping(reqId string, d2 *models.Demo) error {
	// core.Db(consts.DB_DEMO).AutoMigrate(
	// 	models.Demo{},
	// )

	if err := core.Db(consts.DB_DEMO).Create(&d2).Error; err != nil {
		core.Log.Error("REQID:"+reqId, zap.Error(err))
		return err
	}
	return nil
}
