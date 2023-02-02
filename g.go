package wahing

import (
	"github.com/wahing/gorm-generate/conf"
	"github.com/wahing/gorm-generate/dbtools"
	"github.com/wahing/gorm-generate/generate"
)

// "gorm-generate/dbtools"
// "gorm-generate/generate"

func Gen(db conf.DbConf, table string) {
	// fmt.Println("1")
	//初始化数据库
	dbtools.Init(db)
	//generate.Generate() //生成所有表信息
	generate.Generate(table) //生成指定表信息，可变参数可传入多个表名
}
