package redis

import (
	"github.com/gomodule/redigo/redis"
	"ws/config"
)

var Pool = newPool()

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.C.GetString("redis.host"))
			if _, err := c.Do("AUTH", config.C.GetString("redis.password")); err != nil {
				c.Close()
				return nil, err
			}
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}
