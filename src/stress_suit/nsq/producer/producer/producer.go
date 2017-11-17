package producer

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	nsq "github.com/nsqio/go-nsq"
)

var wg sync.WaitGroup
var stop bool
var lastCount int
var w *nsq.Producer

func StartProduce() {
	config := nsq.NewConfig()
	w, _ = nsq.NewProducer("127.0.0.1:4150", config)
	stop = false

	count := 1000
	lastCount = count

	wg.Add(count)

	for index := 0; index < count; index++ {
		topic := "nsq_crazy_topic_" + strconv.Itoa(index)
		fmt.Printf("topic: %s /n", topic)

		go func() {
			defer wg.Done()
			for {
				if stop {
					break
				}
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
}

func MadProducer(topicNumber int) {
	if topicNumber <= lastCount {
		return
	}

	tmpCount := lastCount
	lastCount = topicNumber

	for index := tmpCount; index < topicNumber; index++ {
		topic := "nsq_crazy_topic_" + strconv.Itoa(index)
		fmt.Printf("topic: %s /n", topic)

		go func() {
			for {
				if stop {
					break
				}
				err := w.Publish(topic, []byte("test"))
				if err != nil {
					log.Panic("Could not connect")
				}
				time.Sleep(time.Second * 2)
			}
		}()
	}
}

func Stop() {
	stop = true
}
