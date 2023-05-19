package db

import (
	"context"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/go_web_scaffold/model"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var db *gorm.DB

func init() {
	dbConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: util.NewGormLog(util.NewDefaultGormSqlHandle()),
	}

	var err error
	db, err = initDb(dbConfig)
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

func initDb(dbConfig *gorm.Config) (*gorm.DB, error) {
	err := util.CreateFolderPath(util.GenCtx(), model.ResourcePath)
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(sqlite.Open(fmt.S), dbConfig)
	if err != nil {
		return db, err
	}
	err = db.AutoMigrate(&model.Config{})
	return db, err
}

func getDb(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}

func getWhere(ctx context.Context, where *gorm.DB) *gorm.DB {
	if where != nil {
		return where
	}
	return getDb(ctx)
}
