package service

import (
	"fmt"

	"github.com/hugplus/go-walker/common/consts"
	"github.com/hugplus/go-walker/common/errs"
	"github.com/hugplus/go-walker/common/errs/codes"
	"github.com/hugplus/go-walker/core"
	"github.com/hugplus/go-walker/modules/demo/models"
	"github.com/hugplus/go-walker/modules/demo/service/dto"
	"go.uber.org/zap"
)

type DemoService struct {
	//base.BaseService
}

func (e *DemoService) Page(req dto.DemePageReq, list *[]models.Demo, total *int64, reqId string) errs.IError {
	db := core.Db(consts.DB_DEMO)
	//TODO WHERE
	fmt.Printf("size:%d offset:%d", req.GetSize(), req.GetOffset())
	if err := db.Limit(req.GetSize()).Offset(req.GetOffset()).Find(list).Count(total).Error; err != nil {
		berr := errs.Err(codes.FAILURE, reqId, err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	return nil
}

func (*DemoService) Create(data *models.Demo, reqId string) errs.IError {
	if err := core.Db(consts.DB_DEMO).Create(data).Error; err != nil {
		berr := errs.Err(codes.FAILURE, reqId, err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	return nil
}

func (*DemoService) Update(data *models.Demo, reqId string) errs.IError {
	if err := core.Db(consts.DB_DEMO).Save(data).Error; err != nil {
		berr := errs.Err(codes.FAILURE, reqId, err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	return nil
}

func (*DemoService) Get(id int, data *models.Demo, reqId string) errs.IError {
	if err := core.Db(consts.DB_DEMO).First(data, id).Error; err != nil {
		berr := errs.Err(codes.FAILURE, reqId, err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	return nil
}
