// description: 用于初始化 gorm.DB
//
// author: vignetting
// time: 2021/5/10

package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"structure/pkg/setting"
)

var db *gorm.DB

func SetUp() {
	var err error
	if db, err = gorm.Open(mysql.Open(setting.DatabaseSetting.Url), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,                                       // 不创建外键
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true}, // 不让表名自动加 s
	}); err != nil {
		panic("数据库连接创建失败，" + err.Error())
	} else {
		if sqlDB, err := db.DB(); err != nil {
			panic("获取 sqlDB 失败，失败原因为：" + err.Error())
		} else {
			// 设置最大空闲连接数
			sqlDB.SetMaxIdleConns(setting.DatabaseSetting.MaxIdleConnection)

			// 设置最大连接数
			sqlDB.SetMaxOpenConns(setting.DatabaseSetting.MaxOpenConnection)

			// 设置连接的最大空闲时间
			sqlDB.SetConnMaxIdleTime(setting.DatabaseSetting.MaxIdleTime)

			// 设置连接的最大复用时间
			sqlDB.SetConnMaxLifetime(setting.DatabaseSetting.MaxLifeTime)
		}
	}
}
