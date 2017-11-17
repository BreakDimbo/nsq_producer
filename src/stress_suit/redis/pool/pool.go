package pool

import (
	"flag"
	"fmt"
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
	c := pool.Get()
	_, err := c.Do("AUTH", "ORjPtnqVDlrlnkP5KoT5")
	if err != nil {
		fmt.Printf("failed to auth redis. cache disabled.err:%s", err.Error())
		return nil
	}
	return c
}

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}
