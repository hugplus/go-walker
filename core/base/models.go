package base

import (
	"time"

	"gorm.io/gorm"
)

type ControlBy struct {
	CreateBy int `json:"createBy" gorm:"index;comment:创建者"` //创建者id
	UpdateBy int `json:"updateBy" gorm:"index;comment:更新者"` //更新者id
}

type Model struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"` //主键
}

type ModelIntTime struct {
	CreatedAt int `json:"createdAt" gorm:"comment:创建时间"`   //创建时间戳
	UpdatedAt int `json:"updatedAt" gorm:"comment:最后更新时间"` //更新时间戳
	//DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}

type ModelTime struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`   //创建时间
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:最后更新时间"` //更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`     //删除时间
}
