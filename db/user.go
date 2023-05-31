package db

import (
	"context"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/go_web_scaffold/model"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserService struct {
	util.GormService[model.User, model.UserInquiry]
}

var User = UserService{GormService: util.GormService[model.User, model.UserInquiry]{GormHandler: &UserHandler{}}}

func (this *UserService) Add(ctx context.Context, object ...*model.User) ([]*model.User, error) {
	for i := range object {
		if object[i].Username == "" {
			logrus.WithContext(ctx).WithFields(logrus.Fields{}).Warn("保存用户，Username为空")
			return object, errors.Errorf("保存用户，Username为空")
		}
		if object[i].Password == "" {
			logrus.WithContext(ctx).WithFields(logrus.Fields{}).Warn("保存用户，Password为空")
			return object, errors.Errorf("保存用户，Password为空")
		}
	}
	return this.Insert(ctx, object...)
}
func (this *UserService) Remove(ctx context.Context, inquiry model.UserInquiry) error {
	return this.Delete(ctx, inquiry)
}
func (this *UserService) Edit(ctx context.Context, object *model.User) (*model.User, int64, error) {
	if object.Username == "" {
		logrus.WithContext(ctx).WithFields(logrus.Fields{}).Warn("编辑用户，Username为空")
		return object, 0, errors.Errorf("编辑用户，Username为空")
	}
	if object.Password == "" {
		logrus.WithContext(ctx).WithFields(logrus.Fields{}).Warn("编辑用户，Password为空")
		return object, 0, errors.Errorf("编辑用户，Password为空")
	}
	return this.Update(ctx, object)
}
func (this *UserService) List(ctx context.Context, inquiry model.UserInquiry) ([]*model.User, int64, error) {
	return this.Select(ctx, inquiry)
}
func (this *UserService) Get(ctx context.Context, inquiry model.UserInquiry) (*model.User, error) {
	return this.SelectOne(ctx, inquiry)
}

type UserHandler struct {
}

func (this *UserHandler) GetName(ctx context.Context) string {
	return "用户"
}
func (this *UserHandler) GetDb(ctx context.Context, where *gorm.DB) *gorm.DB {
	return getDb(ctx, where)
}
func (this *UserHandler) Where(ctx context.Context, where *gorm.DB, inquiry model.UserInquiry) *gorm.DB {
	if inquiry.Id > 0 {
		where = getDb(ctx, where).Where("id = ?", inquiry.Id)
	}
	if inquiry.Username != "" {
		where = getDb(ctx, where).Where("user_name = ?", inquiry.Username)
	}
	if inquiry.Password != "" {
		where = getDb(ctx, where).Where("password = ?", inquiry.Password)
	}
	return where
}
