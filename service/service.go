package service

import (
	"context"
	"github.com/cellargalaxy/go_web_scaffold/config"
	"github.com/cellargalaxy/go_web_scaffold/model"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func Login(ctx context.Context, username, password string) (model.Config, error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{"username": username, "password": password}).Info("登录")
	if username != config.Config.Username || password != config.Config.Password {
		logrus.WithContext(ctx).WithFields(logrus.Fields{}).Error("登录失败")
		return model.Config{}, errors.Errorf("登录失败")
	}
	return config.Config, nil
}
