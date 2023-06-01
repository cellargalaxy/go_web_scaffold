package model

import (
	"github.com/cellargalaxy/go_common/util"
)

type Config struct {
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

func (this Config) String() string {
	return util.JsonStruct2String(this)
}
