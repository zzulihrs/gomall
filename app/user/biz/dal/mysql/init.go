package mysql

import (
	"fmt"
	"github.com/cloudwego/biz-demo/gomall/app/user/conf"
	"github.com/cloudwego/biz-demo/gomall/app/user/model"
	"github.com/cloudwego/biz-demo/gomall/common/mtl"
	"gorm.io/plugin/opentelemetry/tracing"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	//fmt.Println("dsn:" + dsn)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics(), tracing.WithTracerProvider(mtl.TracerProvider))); err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.User{}) // 数据库自动迁移
	if err != nil {
		panic(err)
	}
}
