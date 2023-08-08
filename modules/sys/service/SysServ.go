package service

import (
	"fmt"

	"github.com/hugplus/go-walker/core"
	"github.com/hugplus/go-walker/modules/sys/models"
	"github.com/hugplus/go-walker/modules/sys/service/dto"
)

type SysServ struct {
}

func (s SysServ) PingS(c *dto.SysDto) error {

	// core.Db("master").AutoMigrate(
	// 	models.Sys{},
	// )

	d := models.Sys{
		Name: (*c).Name,
	}
	fmt.Println(d)

	cstr, err := core.Cache.Get("test")
	fmt.Printf("data :%s ,err:%v\n", cstr, err)

	if err := core.Db("master").Create(&d).Error; err != nil {
		return err
	}
	// if err := core.Cache.Set("test", 1, time.Hour); err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	var d2 models.Sys

	if err := core.Db("master").First(&d2, 1).Error; err != nil {
		return err
	}

	fmt.Printf("data :%v \n", d2)

	return nil
}
