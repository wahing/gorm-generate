package conf

// model保存路径
const ModelPath = "./internal/model/"

// 是否覆盖已存在model
const ModelReplace = true

// 数据库驱动
const DriverName = "mysql"

type DbConf struct {
	Host   string
	Port   string
	User   string
	Pwd    string
	DbName string
}

// 数据库链接配置
var MasterDbConfig DbConf = DbConf{
	Host: "",
	Port: "",
	User: "",
	Pwd:  "",

	DbName: "",
}
