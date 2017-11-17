package main

import (
	"fmt"
	"log"
	"net/http"
)

// ListenAndServe 监听 APP 的请求并处理
func ListenAndServe() {
	router := NewRouter()
	host := "127.0.0.1"
	port := 5688

	addr := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf("addr %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
