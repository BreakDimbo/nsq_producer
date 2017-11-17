package pool

import (
	"flag"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	pool        *redis.Pool
	redisServer = flag.String("redisServer", ":6379", "")
)

func Init(addr string) {
	pool = newPool(addr)
}

func GetConn() redis.Conn {
	return pool.Get()
}

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}
