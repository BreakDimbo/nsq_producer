package main

import (
	"stress_suit/redis/pool"
	"stress_suit/redis/pub/publisher"
)

func main() {
	pool.Init("127.0.0.1:6704")
	publisher.Init()
	go publisher.StartPub()
	ListenAndServe()
}
