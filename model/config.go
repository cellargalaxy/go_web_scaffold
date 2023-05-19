package model

import "github.com/cellargalaxy/go_common/model"

type Config struct {
	model.Model
	Username string `json:"user_name" yaml:"user_name"`
	Password string `json:"password" yaml:"password"`
}
