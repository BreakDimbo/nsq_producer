package main

import (
	"stress_suit/redis/pool"
	"stress_suit/redis/sub/subscribe"
)

func main() {
	pool.Init("10.30.248.116:6704")
	subscribe.Init()
	subscribe.StartSub()
	ListenAndServe()
}
