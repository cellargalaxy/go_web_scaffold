package model

import (
	"github.com/cellargalaxy/go_common/util"
)

type Config struct {
	Username string `json:"user_name" yaml:"user_name"`
	Password string `json:"password" yaml:"password"`
}

func (this Config) String() string {
	return util.JsonStruct2String(this)
}
