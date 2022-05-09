package driver

import (
	"fmt"
	"redigo-example/config"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/labstack/gommon/log"
)

func InitRedis(redisConfig config.RedisConfig) (pool *redis.Pool) {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port))
			if err != nil {
				return nil, err
			}

			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxActive: redisConfig.MaxActive,
		MaxIdle:   redisConfig.MaxIdle,
	}

	return
}

func Ping(redisDB *redis.Pool) error {

	conn := redisDB.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}

	log.Info("success connect redis!")
	return nil
}
