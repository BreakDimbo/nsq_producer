package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	nsq "github.com/nsqio/go-nsq"
)

type subConfig struct {
	nsqConfig        *nsq.Config
	nsqLookUpAddress string
	nsqdHTTPAddress  string
}

var subCfg *subConfig
var nsqChannelName = "crazy_consumer"
var consumers map[string]*nsq.Consumer

// InitSub 初始化 Consumer 的配置
func InitSub() {
	cfg := nsq.NewConfig()

	lookupHost := "47.93.79.149"
	lookupPort := 4161
	nsqLookUpAddress := fmt.Sprintf("%s:%d", lookupHost, lookupPort)
	nsqdHost := "47.93.79.149"
	nsqdHTTPPort := "4151"
	nsqdHTTPAddress := fmt.Sprintf("%s:%s", nsqdHost, nsqdHTTPPort)
	subCfg = &subConfig{
		nsqConfig: cfg,
		// TODO: 这里需要使用数组
		nsqLookUpAddress: nsqLookUpAddress,
		nsqdHTTPAddress:  nsqdHTTPAddress,
	}
}

func StartConsumers() {
	count := 1000
	for index := 0; index < count; index++ {
		topic := "nsq_crazy_topic_" + strconv.Itoa(index)
		fmt.Printf("topic: %s /n", topic)

		c := NewConsumer(topic)
		consumers[topic] = c
	}
}

// NewConsumer 创建一个新的消费者
func NewConsumer(topic string) *nsq.Consumer {
	nsq.Register(topic, nsqChannelName)
	q, _ := nsq.NewConsumer(topic, nsqChannelName, subCfg.nsqConfig)
	q.AddHandler(nsq.HandlerFunc(nsqHandler))
	err := q.ConnectToNSQLookupd(subCfg.nsqLookUpAddress)
	if err != nil {
		fmt.Errorf("Could not connect")
	}
	return q
}

func nsqHandler(message *nsq.Message) error {
	fmt.Printf("[NSQ Message GET]: %s", message.Body)
	return nil
}

func deleteChannel(topic string) {
	topicStr := url.QueryEscape(topic)
	urlStr := fmt.Sprintf("http://%s/channel/delete?topic=%s&channel=%s", subCfg.nsqdHTTPAddress, topicStr, nsqChannelName)

	fmt.Printf("delete nsq channel url: %s", urlStr)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, nil)
	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf("delete nsq channel error: %s", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("delete nsq channel result: %s", resp)

}

// 保证消费者断开连接后，channel 会被销毁
func ephemeralStr(ch string) string {
	return ch
}
