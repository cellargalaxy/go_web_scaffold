package model

import (
	"github.com/cellargalaxy/go_common/model"
	"github.com/cellargalaxy/go_common/util"
)

type User struct {
	model.Model
	Username string `json:"user_name" yaml:"user_name" gorm:"user_name"`
	Password string `json:"password" yaml:"password" gorm:"password"`
}

func (this User) String() string {
	return util.JsonStruct2String(this)
}
func (this User) TableName() string {
	return "user"
}

type UserInquiry struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	User
}

func (this UserInquiry) String() string {
	return util.JsonStruct2String(this)
}
func (this UserInquiry) GetOffset() int {
	return this.Offset
}
func (this UserInquiry) GetLimit() int {
	return this.Limit
}
