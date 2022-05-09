package repository

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type RedisRepositoryIn interface {
	CheckExist(key string) (exist bool, err error)
	GetCache(key string) (data []byte, err error)
	WriteCache(key string, data []byte) (err error)
	DeleteCache(key string) (err error)
}

type redisRepository struct {
	redisPool *redis.Pool
}

func NewRedisRepository(redisPool *redis.Pool) *redisRepository {
	return &redisRepository{redisPool}
}

func (r *redisRepository) CheckExist(key string) (exist bool, err error) {
	conn := r.redisPool.Get()
	defer conn.Close()

	exist, err = redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return exist, fmt.Errorf("%s doesn't exist in redis", key)
	}

	return
}

func (r *redisRepository) GetCache(key string) (data []byte, err error) {
	conn := r.redisPool.Get()
	defer conn.Close()

	data, err = redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, fmt.Errorf("error getting data on key %s", key)
	}

	return
}

func (r *redisRepository) WriteCache(key string, data []byte) (err error) {
	conn := r.redisPool.Get()
	defer conn.Close()

	_, err = conn.Do("SET", key, data)
	if err != nil {
		return fmt.Errorf("error write data on key %s", key)
	}

	return
}

func (r *redisRepository) DeleteCache(key string) (err error) {
	conn := r.redisPool.Get()
	defer conn.Close()

	_, err = conn.Do("DEL", key)
	if err != nil {
		return fmt.Errorf("error delete data on key %s", key)
	}

	return
}
