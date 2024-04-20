package global

type Config struct {
	MysqlConfig  MySqlConfig  `ini:"mysql"`
	ServerConfig ServerConfig `ini:"server"`
}

type MySqlConfig struct {
	Datasource string `ini:"datasource"`
}

type ServerConfig struct {
	ServerPort string `ini:"server_port"`
	AppMode    string `ini:"app_mode"`
}
