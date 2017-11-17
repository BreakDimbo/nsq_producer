package main

import "stress_suit/nsq/producer/producer"

func main() {
	go producer.StartProduce()
	ListenAndServe()
}
