package main

import "stress_suit/nsq/pro/producer"

func main() {
	go producer.StartProduce()
	ListenAndServe()
}
