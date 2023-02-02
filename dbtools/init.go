package dbtools

import (
	"log"
	"os"

	"github.com/wahing/gorm-generate/conf"
)

func Init(db conf.DbConf) {
	//初始化Mysql连接池
	mysql := GetMysqlInstance().InitMysqlPool(db)
	if !mysql {
		log.Println("init database pool failure...")
		os.Exit(1)
	}
}
