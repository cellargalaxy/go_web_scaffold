package model

import (
	"github.com/cellargalaxy/go_common/model"
	"github.com/cellargalaxy/go_common/util"
)

type User struct {
	model.Model
	Username string `json:"username" yaml:"username" gorm:"username"`
	Password string `json:"password" yaml:"password" gorm:"password"`
}

func (this User) String() string {
	return util.JsonStruct2String(this)
}
func (this User) TableName() string {
	return "user"
}

type UserInquiry struct {
	model.Inquiry
	User
}

func (this UserInquiry) String() string {
	return util.JsonStruct2String(this)
}
