package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	nsq "github.com/nsqio/go-nsq"
)

var wg sync.WaitGroup

func main() {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	count := 10000
	wg.Add(count)

	for index := 0; index < count; index++ {
		topic := "nsq_crazy_topic_" + strconv.Itoa(index)
		fmt.Printf("topic: %s", topic)
		fmt.Println()

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

	defer w.Stop()
	wg.Wait()
}
