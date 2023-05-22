package db

import (
	"context"
	"fmt"
	"github.com/cellargalaxy/go_web_scaffold/model"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

func InsertScaffold(ctx context.Context, object ...*model.Scaffold) ([]*model.Scaffold, error) {
	if len(object) == 0 {
		logrus.WithContext(ctx).WithFields(logrus.Fields{}).Warn("插入脚手架，为空")
		return object, nil
	}
	where := getDb(ctx)
	err := where.Create(&object).Error
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("插入脚手架，异常")
		return object, fmt.Errorf("插入脚手架，异常: %+v", err)
	}
	logrus.WithContext(ctx).WithFields(logrus.Fields{}).Info("插入脚手架，完成")
	return object, nil
}

func whereScaffold(ctx context.Context, where *gorm.DB, inquiry model.ScaffoldInquiry) *gorm.DB {
	if inquiry.Id > 0 {
		where = getWhere(ctx, where).Where("id = ?", inquiry.Id)
	}
	if inquiry.Username != "" {
		where = getWhere(ctx, where).Where("user_name = ?", inquiry.Username)
	}
	if inquiry.Password != "" {
		where = getWhere(ctx, where).Where("password = ?", inquiry.Password)
	}
	return where
}

func DeleteScaffold(ctx context.Context, inquiry model.ScaffoldInquiry) error {
	var where *gorm.DB
	where = whereScaffold(ctx, where, inquiry)
	if where == nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"inquiry": inquiry}).Warn("删除脚手架，条件为空")
		return fmt.Errorf("删除脚手架，条件为空")
	}
	err := where.Delete(&inquiry.Scaffold).Error
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Warn("删除脚手架，异常")
		return fmt.Errorf("删除脚手架，异常: %+v", err)
	}
	logrus.WithContext(ctx).WithFields(logrus.Fields{}).Info("删除脚手架，完成")
	return nil
}

func UpdateScaffold(ctx context.Context, object *model.Scaffold) (*model.Scaffold, error) {
	if object == nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{}).Warn("更新脚手架，为空")
		return object, nil
	}
	object.UpdatedAt = time.Now()
	where := getDb(ctx)
	where = where.Model(&object)
	where = where.Where("id = ?", object.Id)
	err := where.Select("*").Updates(&object).Error
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Warn("更新脚手架，异常")
		return object, errors.Errorf("更新脚手架，异常: %+v", err)
	}
	logrus.WithContext(ctx).WithFields(logrus.Fields{}).Info("更新脚手架，完成")
	return object, nil
}
