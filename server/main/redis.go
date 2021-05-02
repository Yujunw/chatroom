package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

// 定义一个全局的pool
var pool *redis.Pool

func initPool(addr string, maxIdle, maxActive int, idleTime time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTime,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
}
