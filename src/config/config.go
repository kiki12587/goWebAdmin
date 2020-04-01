package config

type Env struct {
	Static             string // 静态文件相对路径
	ServerPort         string // web服务端口号
	DatabaseAddr       string //数据库地址
	DatabaseProt       string //端口
	DatabaseUsername   string //账号
	DatabasePassword   string //密码
	DatabaseTablename  string //数据表
	DatabaseMaxIdleCon int    //最大空闲连接数
	DatabaseMaxOpenCon int    //最大连接数
}

func GetEnv() *Env {
	return &env
}
