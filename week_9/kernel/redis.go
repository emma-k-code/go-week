package kernel

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	Conn redis.Conn
}

func (r *Redis) CreateConn() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Printf("Connect to redis error: %+v\n", err)
	}

	r.Conn = c
}

func (r *Redis) Set(key string, value string) bool {
	_, err := r.Conn.Do("SET", key, value)
	if err != nil {
		fmt.Printf("redis set failed: %+v\n", err)
		return false
	}

	return true
}

func (r *Redis) Get(key string) string {
	result, err := redis.String(r.Conn.Do("GET", key))
	if err != nil {
		fmt.Printf("redis get failed: %+v\n", err)
		return ""
	}

	return result
}
