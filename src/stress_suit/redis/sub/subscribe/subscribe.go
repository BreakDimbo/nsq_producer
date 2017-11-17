package subscribe

import (
	"fmt"
	"strconv"
	"stress_suit/redis/pool"

	"github.com/garyburd/redigo/redis"
)

var stop bool
var cCount int

func Init() {
	stop = false
	cCount = 1000
}

func StartSub() {
	for index := 0; index < cCount; index++ {
		topic := "nsq_crazy_topic_" + strconv.Itoa(index)
		fmt.Printf("topic: %s /n", topic)

		go sub(topic)
	}
}

func UpdateCount(count int) {
	if count <= cCount {
		return
	}

	for index := cCount; index < count; index++ {
		topic := "nsq_crazy_topic_" + strconv.Itoa(index)
		fmt.Printf("topic: %s /n", topic)

		go sub(topic)
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

		go sub(topic)
	}
	cCount = endCount
}

func Stop() {
	stop = true
}

func sub(topic string) {
	for !stop {
		// Get a connection from a pool
		c := pool.GetConn()
		psc := redis.PubSubConn{Conn: c}

		// Set up subscriptions
		err := psc.Subscribe(topic)
		if err != nil {
			fmt.Printf("Subscribe Error: %s", err)
			return
		}

		// While not a permanent error on the connection.
		for c.Err() == nil && !stop {
			switch v := psc.Receive().(type) {
			case redis.Message:
				fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
			case redis.Subscription:
				fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
			case error:
				fmt.Print("receive Error")
				return
			}
		}
		c.Close()
	}
}
