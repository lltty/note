package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func hello(w http.ResponseWriter, res *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("hello[%s]", params.ByName("name"))))
}

//统计执行耗时
func StatisticsTimeMiddle(next httprouter.Handle) httprouter.Handle {

	return httprouter.Handle(func(w http.ResponseWriter, res *http.Request, params httprouter.Params) {
		start := time.Now()
		next(w, res, params)
		timeSlapsed := time.Since(start)
		log.Printf("用户[%v]%s-请求[%s],耗时:%v\n", params.ByName("name"), res.Method, res.URL.Path, timeSlapsed)
	})
}

//统计调用次数
func CountReq(next httprouter.Handle) httprouter.Handle {

	s := make(map[string]uint)
	return httprouter.Handle(func(w http.ResponseWriter, res *http.Request, params httprouter.Params) {
		path := res.URL.Path
		s[path] += 1
		next(w, res, params)
		log.Printf("用户[%v]%s-请求[%s],次数:%v\n", params.ByName("name"), res.Method, path, s[path])
	})
}

func main() {

	port := ":9001"
	r := httprouter.New()
	//这边封装了多个中间件
	r.GET("/hello/:name", CountReq(StatisticsTimeMiddle(hello)))
	http.ListenAndServe(port, r)

}
