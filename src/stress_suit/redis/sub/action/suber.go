package action

import (
	"fmt"
	"net/http"
	"strconv"
	"stress_suit/redis/sub/subscribe"
)

func Mad(res http.ResponseWriter, req *http.Request) {
	count := queryParam(req, "count")
	topicCount, _ := strconv.Atoi(count)
	fmt.Printf("Mad producer with count: %d", topicCount)
	subscribe.UpdateCount(topicCount)
}

func Stop(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Stop")
	subscribe.Stop()
}

func Set(res http.ResponseWriter, req *http.Request) {
	start := queryParam(req, "start")
	end := queryParam(req, "end")
	startCount, _ := strconv.Atoi(start)
	endCount, _ := strconv.Atoi(end)
	fmt.Printf("Mad producer with count: %d", startCount)
	fmt.Printf("Mad producer with count: %d", endCount)
	subscribe.SetCount(startCount, endCount)
}

func queryParam(req *http.Request, key string) string {
	return req.URL.Query().Get(key)
}
