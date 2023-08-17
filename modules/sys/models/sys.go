package models

import (
	"encoding/json"

	"github.com/hugplus/go-walker/core/base"
)

type Sys struct {
	base.Model
	Name string `json:"name" gorm:"type:varchar(128);comment:名字"`
	base.StatusModel
	base.ControlBy
	base.ModelIntTime
}

func (m Sys) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}
