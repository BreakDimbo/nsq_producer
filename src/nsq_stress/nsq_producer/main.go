package main

import "nsq_stress/nsq_producer/producer"

func main() {
	go producer.StartProduce()
	ListenAndServe()
}
