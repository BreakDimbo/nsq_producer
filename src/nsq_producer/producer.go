package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	nsq "github.com/nsqio/go-nsq"
)

var consumers map[string]*nsq.Consumer
var wg sync.WaitGroup

func StartProduce() {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)
	consumers = make(map[string]*nsq.Consumer)

	count := 1000
	wg.Add(count)

	for index := 0; index < count; index++ {
		topic := "nsq_crazy_topic_" + strconv.Itoa(index)
		fmt.Printf("topic: %s /n", topic)

		c := NewConsumer(topic)
		consumers[topic] = c

		go func() {
			defer wg.Done()
			for {
				err := w.Publish(topic, []byte("test"))
				if err != nil {
					log.Panic("Could not connect")
				}
				time.Sleep(time.Second * 2)
			}
		}()
	}

	wg.Wait()
	defer w.Stop()
	defer func() {
		for k, v := range consumers {
			v.Stop()
			deleteChannel(k)
		}
	}()
}
