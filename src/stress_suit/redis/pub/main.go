package main

import (
	"stress_suit/redis/pool"
	"stress_suit/redis/pub/publisher"
)

func main() {
	pool.Init("47.93.79.149:6704")
	publisher.Init()
	go publisher.StartPub()
	ListenAndServe()
}
