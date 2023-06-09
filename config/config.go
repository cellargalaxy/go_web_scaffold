package config

import (
	"context"
	"fmt"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/go_web_scaffold/model"
	"path"
)

var Config = model.Config{
	Username: "admin",
	Password: "123456",
}

func init() {
	ctx := util.GenCtx()
	service := util.NewConfigService(&ConfigHandler{})
	err := service.Start(ctx)
	if err != nil {
		panic(err)
	}
}

type ConfigHandler struct {
}

func (this *ConfigHandler) GetPath(ctx context.Context) string {
	return path.Join(model.ResourcePath, fmt.Sprintf("%s.yaml", model.ServerName))
}

func (this *ConfigHandler) GetConfig(ctx context.Context) string {
	return util.YamlStruct2String(ctx, Config)
}

func (this *ConfigHandler) ParseConfig(ctx context.Context, text string) error {
	var config model.Config
	err := util.YamlString2Struct(ctx, text, &config)
	if err != nil {
		return err
	}
	Config = config
	return nil
}
