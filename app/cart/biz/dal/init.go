package dal

import (
	"github.com/cloudwego/biz-demo/gomall/app/cart/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
