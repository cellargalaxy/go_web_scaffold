package main

import (
	"fmt"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/go_web_scaffold/model"
	"github.com/sirupsen/logrus"
)

func init() {
	util.Init(model.ServerName)
}

func main() {
	ctx := util.GenCtx()
	fmt.Println(util.GetLogIdString(ctx))
	logrus.WithContext(ctx).WithFields(logrus.Fields{}).Info("打印日志")
}
