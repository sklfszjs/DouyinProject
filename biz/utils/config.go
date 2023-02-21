package utils

type ServerConfig struct {
	DBUsrName  string
	DBPassword string
	DBName     string
	IP         string
	Port       int
}

var Configs *ServerConfig

func init() {
	Configs = &ServerConfig{
		DBUsrName:  "root",
		DBPassword: "#Aa2101572..",
		DBName:     "douyin",
		IP:         "127.0.0.1",
		Port:       8888,
	}
}
func GetConfigs() *ServerConfig {
	return Configs
}
