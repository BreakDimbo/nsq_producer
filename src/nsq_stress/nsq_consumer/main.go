package main

import "time"

func main() {
	InitSub()
	StartConsumers()
	time.Sleep(1 * time.Hour)
}
