package config

const (
	Host     = "127.0.0.1"
	Port     = "5432"
	User     = "root"
	Password = "1234"
	DbName   = "postgres"
	SslMode  = "disable"
)

type DataBaseConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	DbName   string
	SslMode  string
}

type Config struct {
	DbConfig DataBaseConfig
}

func NewConfig() *Config {
	return &Config{
		DbConfig: DataBaseConfig{
			Host:     Host,
			User:     User,
			Password: Password,
			Port:     Port,
			DbName:   DbName,
			SslMode:  SslMode,
		}}
}
