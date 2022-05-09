package config

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Host string
	Port int
}

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

type RedisConfig struct {
	Host      string
	Port      int
	Password  string
	MaxActive int
	MaxIdle   int
	UserKey   string
	BookKey   string
}

func InitConfig() (appConf AppConfig, dbConf DBConfig, redisConf RedisConfig) {
	v := viper.New()
	v.AddConfigPath("./config")
	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err == nil {
		log.Info("Using config file: ", v.ConfigFileUsed())
	} else {
		log.Panic(fmt.Errorf("config error: %s", err))
	}

	appConf.Host = v.GetString("app.host")
	appConf.Port = v.GetInt("app.port")

	dbConf.Host = v.GetString("database.host")
	dbConf.Port = v.GetInt("database.port")
	dbConf.Username = v.GetString("database.username")
	dbConf.Password = v.GetString("database.password")
	dbConf.Name = v.GetString("database.name")

	redisConf.Host = v.GetString("redis.host")
	redisConf.Port = v.GetInt("redis.port")
	redisConf.Password = v.GetString("redis.password")
	redisConf.MaxActive = v.GetInt("redis.max_active")
	redisConf.MaxIdle = v.GetInt("redis.max_idle")
	redisConf.UserKey = v.GetString("redis.user_key")
	redisConf.BookKey = v.GetString("redis.book_key")

	return appConf, dbConf, redisConf
}
