package model

import "github.com/cellargalaxy/go_common/util"

type Scaffold struct {
	Config
}

func (this Scaffold) String() string {
	return util.JsonStruct2String(this)
}

type ScaffoldInquiry struct {
	Scaffold
}

func (this ScaffoldInquiry) String() string {
	return util.JsonStruct2String(this)
}
