package publisher

import (
	"fmt"
	"strconv"
	"stress_suit/redis/pool"
	"time"
)

var stop bool
var cCount int

func Init() {
	stop = false
	cCount = 1000
}

func StartPub() {

	for index := 0; index < cCount; index++ {
		topic := "nsq_crazy_topic_" + strconv.Itoa(index)
		fmt.Printf("topic: %s /n", topic)

		go pub(topic, "crazy Message!")
	}
}

func UpdateCount(count int) {
	if count <= cCount {
		return
	}

	for index := cCount; index < count; index++ {
		topic := "nsq_crazy_topic_" + strconv.Itoa(index)
		fmt.Printf("topic: %s /n", topic)

		go pub(topic, "crazy Message!")
	}
	cCount = count
}

func SetCount(startCount int, endCount int) {
	if endCount <= cCount || startCount <= cCount {
		return
	}

	for index := startCount; index < endCount; index++ {
		topic := "nsq_crazy_topic_" + strconv.Itoa(index)
		fmt.Printf("topic: %s /n", topic)

		go pub(topic, "crazy Message!")
	}
	cCount = endCount
}

func Stop() {
	stop = true
}

func pub(topic string, message string) {
	// Get a connection from a pool
	c := pool.GetConn()
	defer c.Close()

	for !stop {
		_, err := c.Do("PUBLISH", topic, message)
		if err != nil {
			fmt.Printf("Send Message error: %s", err)
		}
		time.Sleep(3000 * time.Millisecond)
	}
}
