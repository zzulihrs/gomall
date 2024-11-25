package mysql

import (
	"fmt"
	"github.com/cloudwego/biz-demo/gomall/app/order/biz/model"
	"github.com/cloudwego/biz-demo/gomall/app/order/conf"
	"github.com/cloudwego/biz-demo/gomall/common/mtl"
	"github.com/cloudwego/kitex/pkg/klog"
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

	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics(), tracing.WithTracerProvider(mtl.TracerProvider))); err != nil {
		panic(err)
	}

	if os.Getenv("GO_ENV") != "online" {
		err := DB.AutoMigrate(&model.Order{}, &model.OrderItem{})
		if err != nil {
			klog.Error(err)
		}
	}
	if err != nil {
		panic(err)
	}
}
