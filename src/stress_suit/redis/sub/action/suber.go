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

func queryParam(req *http.Request, key string) string {
	return req.URL.Query().Get(key)
}
