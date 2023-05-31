package main

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/go_web_scaffold/controller"
	"github.com/cellargalaxy/go_web_scaffold/db"
	"github.com/cellargalaxy/go_web_scaffold/model"
	"github.com/sirupsen/logrus"
)

func init() {
	util.Init(model.ServerName)
}

func main() {
	ctx := util.GenCtx()
	logrus.WithContext(ctx).WithFields(logrus.Fields{"version": model.Version}).Info()
	db.Init(ctx)
	controller.Init()
}
