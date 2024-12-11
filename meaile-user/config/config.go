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
type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	Database int    `mapstructure:"database" json:"database"`
}
type MinioConfig struct {
	EndPoint   string `mapstructure:"end-point" json:"end-point"`
	AccessKey  string `mapstructure:"access-key" json:"access-key"`
	SecretKey  string `mapstructure:"secret-key" json:"secret-key"`
	BucketName string `mapstructure:"bucket-name" json:"bucket-name"`
}
type ServerConfig struct {
	Name        string      `mapstructure:"name" json:"name"`
	Port        int32       `mapstructure:"port" json:"port"`
	MysqlConfig MysqlConfig `mapstructure:"mysql" json:"mysql"`
	JWTConfig   JWTConfig   `mapstructure:"jwt" json:"jwt"`
	RedisConfig RedisConfig `mapstructure:"redis" json:"redis"`
	MinioConfig MinioConfig `mapstructure:"minio" json:"minio"`
}
