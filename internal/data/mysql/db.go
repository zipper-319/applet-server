package mysql

import (
	"applet-server/internal/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func NewDataDB(c *conf.Data) *gorm.DB {
	return GenDB(c.Database.Source)
}

func GenDB(source string) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s?charset=utf8&parseTime=True&loc=Local", source), // DSN data source name
		DefaultStringSize:         255,                                                             // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                            // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                            // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                            // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                           // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(500)
	sqlDB.SetConnMaxIdleTime(time.Hour)

	return db
}

type PageLimitPara struct {
	Offset int
	Count  int
}

type IdLimitPara struct {
	LastId int64
	Count  int
}
