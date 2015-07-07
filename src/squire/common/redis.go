package common

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var RedisPool *redis.Pool

func init() {
	RedisPool = newPool()
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func RedisString(reply interface{}, err error) (string, error) {
	if r, e := redis.String(reply, err); e == nil || e == redis.ErrNil {
		return r, nil
	} else {
		return "", e
	}
}
