package action

import (
	"fmt"
	"net/http"
	"nsq_stress/nsq_producer/producer"
	"strconv"
)

func Mad(res http.ResponseWriter, req *http.Request) {
	topicNumber := queryParam(req, "topic_number")
	topicCount, _ := strconv.Atoi(topicNumber)
	fmt.Printf("Mad producer with count: %d", topicCount)
	producer.MadProducer(topicCount)
}

func Stop(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Stop")
	producer.Stop()
}

func queryParam(req *http.Request, key string) string {
	return req.URL.Query().Get(key)
}
