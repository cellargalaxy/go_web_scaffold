package main

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/go_web_scaffold/controller"
	"github.com/cellargalaxy/go_web_scaffold/db"
	"github.com/cellargalaxy/go_web_scaffold/model"
)

func init() {
	util.Init(model.ServerName)
}

func main() {
	ctx := util.GenCtx()
	db.Init(ctx)
	controller.Init()
}
