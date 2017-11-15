package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	nsq "github.com/nsqio/go-nsq"
)

func StartProduce() {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	count := 1000

	for index := 0; index < count; index++ {
		topic := "nsq_crazy_topic_" + strconv.Itoa(index)
		fmt.Printf("topic: %s /n", topic)

		NewConsumer(topic)

		go func() {
			for {
				err := w.Publish(topic, []byte("test"))
				if err != nil {
					log.Panic("Could not connect")
				}
				time.Sleep(time.Second * 2)
			}
		}()
	}

	defer w.Stop()
}
