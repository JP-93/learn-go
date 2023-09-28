package config

const (
	Host     = "localhost"
	Port     = "5432"
	User     = "root"
	Password = "1234"
	DbName   = "postgres"
)

type DataBaseConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	DbName   string
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
		}}
}
