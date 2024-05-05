package common

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var instance *Config
var lock = &sync.Mutex{}

type Config struct {
	Host        string
	Port        uint16
	Db_host     string
	Db_port     uint16
	Db_name     string
	Db_username string
	Db_password string
}

func newConfig() (*Config, error) {
	viper.SetConfigFile("config.toml")
	viper.AddConfigPath("./")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		Host:        viper.GetString("host"),
		Port:        viper.GetUint16("port"),
		Db_host:     viper.GetString("db_host"),
		Db_port:     viper.GetUint16("db_port"),
		Db_name:     viper.GetString("db_name"),
		Db_username: viper.GetString("db_username"),
		Db_password: viper.GetString("db_password"),
	}, nil
}

func (c *Config) GetDatabaseConnectionString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Tehran", c.Host, c.Db_username, c.Db_password, c.Db_name, c.Db_port)
}

func GetInstance() (*Config, error) {
	if instance == nil {
		lock.Lock()

		defer lock.Unlock()

		if instance == nil {
			ins, err := newConfig()

			if err != nil {
				return nil, err
			}

			instance = ins
		}
	}

	return instance, nil
}
