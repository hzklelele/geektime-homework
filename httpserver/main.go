package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/golang/glog"
)

func main() {
	flag.Set("v", "5")
	// 执行该行，控制台无日志输出
	// flag.Parse()
	glog.V(3).Info("Starting http server...")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	statusCode := 200
	fmt.Println("entering root handler")
	// 获取、设置系统变量 VERSION
	version := os.Getenv("VERSION")
	fmt.Println(version)
	w.Header().Set("VERSION", version)
	// 设置请求响应头
	for k, v := range r.Header {
		w.Header().Set(k, v[0])
	}
	w.WriteHeader(statusCode)
	glog.V(3).Info("Request host: ", r.RemoteAddr, ", Response code: ", statusCode)
}
