package mysql

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewConfig() *Config {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("cannot read config: %v", err)
	}

	host := viper.GetString("MYSQL_HOSTNAME")
	port := viper.GetString("MYSQL_PORT")
	user := viper.GetString("MYSQL_USER")
	password := viper.GetString("MYSQL_PASSWORD")
	dbName := viper.GetString("MYSQL_DATABASE")

	cfg := &Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
	}

	return cfg
}
