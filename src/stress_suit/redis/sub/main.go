package main

import (
	"stress_suit/redis/pool"
	"stress_suit/redis/sub/subscribe"
)

func main() {
	pool.Init("47.93.79.149:6704")
	subscribe.Init()
	subscribe.StartSub()
	ListenAndServe()
}
