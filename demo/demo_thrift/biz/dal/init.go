package dal

import (
	"gomall/biz/dal/mysql"
	"gomall/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
