package main

import (
	//"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)
var pool *redis.Pool
func initPool(address string, maxidle,maxactive int,idletimeout time.Duration){
	pool = &redis.Pool{
		MaxIdle: maxidle,
		MaxActive: maxactive,
		IdleTimeout: idletimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",address)
		},
	}
}