package model

import (
	"github.com/cellargalaxy/go_common/model"
	"github.com/cellargalaxy/go_common/util"
)

type Config struct {
	model.Model
	Username string `json:"user_name" yaml:"user_name" gorm:"user_name"`
	Password string `json:"password" yaml:"password" gorm:"password"`
}

func (this Config) String() string {
	return util.JsonStruct2String(this)
}
