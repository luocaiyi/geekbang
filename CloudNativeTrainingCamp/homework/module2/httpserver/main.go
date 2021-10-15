package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthzHandler)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// 接收客户端 request，并将 request 中带的 header 写入 response header
	for key, values := range r.Header {
		w.Header().Set(key, strings.Join(values, ";"))
	}
	// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	w.Header().Set("VERSION", os.Getenv("VERSION"))
	// Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	w.WriteHeader(http.StatusOK)
	log.Printf("IP:%s \t StatusCode: %d", r.RemoteAddr, http.StatusOK)
	io.WriteString(w, "Hello HTTP")
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "200")
}
