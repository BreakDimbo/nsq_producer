package action

import (
	"fmt"
	"net/http"
	"strconv"
	"stress_suit/redis/pub/publisher"
)

func Mad(res http.ResponseWriter, req *http.Request) {
	count := queryParam(req, "count")
	topicCount, _ := strconv.Atoi(count)
	fmt.Printf("Mad producer with count: %d", topicCount)
	publisher.UpdateCount(topicCount)
}

func Stop(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Stop")
	publisher.Stop()
}

func Set(res http.ResponseWriter, req *http.Request) {
	start := queryParam(req, "start")
	end := queryParam(req, "end")
	startCount, _ := strconv.Atoi(start)
	endCount, _ := strconv.Atoi(end)
	fmt.Printf("Mad producer with count: %d", startCount)
	fmt.Printf("Mad producer with count: %d", endCount)
	publisher.SetCount(startCount, endCount)
}

func queryParam(req *http.Request, key string) string {
	return req.URL.Query().Get(key)
}
