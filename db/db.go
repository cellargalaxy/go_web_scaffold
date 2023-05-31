package db

import (
	"context"
	"fmt"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/go_web_scaffold/model"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"path"
	"path/filepath"
	"time"
)

var db *gorm.DB

func Init(ctx context.Context) {
	dbConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: util.NewDefaultGormLog(),
	}

	var err error
	db, err = initSqlite(ctx, dbConfig, path.Join(model.ResourcePath, fmt.Sprintf("%s.db", model.ServerName)))
	//db, err = initMysql(ctx, dbConfig, "root:123456@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local&tls=skip-verify")
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(1)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(2)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func initSqlite(ctx context.Context, dbConfig *gorm.Config, filePath string) (*gorm.DB, error) {
	folderPath := filepath.Dir(filePath)
	err := util.CreateFolderPath(ctx, folderPath)
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(sqlite.Open(filePath), dbConfig)
	if err != nil {
		return db, err
	}
	err = db.AutoMigrate(&model.User{})
	return db, err
}

func initMysql(ctx context.Context, dbConfig *gorm.Config, dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), dbConfig)
}

func getDb(ctx context.Context, where *gorm.DB) *gorm.DB {
	if where != nil {
		return where
	}
	return db.WithContext(ctx)
}
