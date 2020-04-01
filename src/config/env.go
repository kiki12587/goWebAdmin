package config

var env = Env{
	ServerPort:         "9000",
	Static:             "src/static",
	DatabaseAddr:       "127.0.0.1",
	DatabaseProt:       "3306",
	DatabaseUsername:   "root",
	DatabasePassword:   "zxcZXC112233",
	DatabaseTablename:  "goadmin",
	DatabaseMaxIdleCon: 50,
	DatabaseMaxOpenCon: 150,
}
