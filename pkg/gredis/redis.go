package gredis

import (
	"encoding/json"
	"gin_log/pkg/setting"
	"github.com/gomodule/redigo/redis"
	"time"
)

var redisConn *redis.Pool

func Setup() {
	redisConn = &redis.Pool{
		MaxIdle: setting.RedisSetting.MaxIdle,  // 最大空闲链接数
		MaxActive: setting.RedisSetting.MaxActive,  // 最大活跃链接数
		IdleTimeout: setting.RedisSetting.IdleTimeout,  // 在这段时间内保持空闲后关闭链接
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				_, err = c.Do("AUTH", setting.RedisSetting.Password)
				if err != nil {
					_ = c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
func Set(key string, data any, t int) error {
	conn := redisConn.Get()
	defer func() {
		_ = conn.Close()
	}()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	_, err = conn.Do("EXPIRE", key, t)
	if err != nil {
		return err
	}
	return nil
}
func Exists(key string) bool {
	conn := redisConn.Get()
	defer func() {
		_ = conn.Close()
	}()

	b, _ := redis.Bool(conn.Do("EXISTS", key))
	return b
}
func Get(key string) ([]byte, error) {
	conn := redisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	replay, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return replay, err
}
func Delete(key string) (bool, error) {
	conn := redisConn.Get()
	defer func() {
		_ = conn.Close()
	}()

	return redis.Bool(conn.Do("DEL", key))
}
func LikeDeletes(key string) error {
	conn := redisConn.Get()
	defer func() {
		_ = conn.Close()
	}()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}
	return nil
}
