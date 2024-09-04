package config

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"name" json:"name"`
	UserName string `mapstructure:"username" json:"username"`
	DBName   string `mapstructure:"dbname" json:"dbname"`
	Password string `mapstructure:"password" json:"password"`
}
type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"host"`
}
type ServerConfig struct {
	Name        string      `mapstructure:"name" json:"name"`
	Port        int32       `mapstructure:"port" json:"port"`
	MysqlConfig MysqlConfig `mapstructure:"mysql" json:"mysql"`
	JWTConfig   JWTConfig   `mapstructure:"jwt" json:"jwt"`
}
