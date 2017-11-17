package main

import (
	"stress_suit/redis/pool"
	"stress_suit/redis/pub/publisher"
)

func main() {
	pool.Init("10.30.248.116:6704")
	publisher.Init()
	go publisher.StartPub()
	ListenAndServe()
}
