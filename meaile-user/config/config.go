package config

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"name" json:"name"`
	User     string `mapstructure:"user" json:"user"`
	DBName   string `mapstructure:"dbname" json:"dbname"`
	Password string `mapstructure:"password" json:"password"`
}

type ServerConfig struct {
	Name        string      `mapstructure:"name" json:"name"`
	MysqlConfig MysqlConfig `mapstructure:"mysql" json:"mysql"`
}
