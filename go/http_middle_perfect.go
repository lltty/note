package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type middleware func(handle httprouter.Handle) httprouter.Handle

type Router struct {
	middlewareChain []middleware
	mux             map[string]httprouter.Handle
	NotFound        http.HandlerFunc
}

// ServeHTTP makes the router implement the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if len(r.mux) > 0 {
		if handler, ok := r.mux[path]; ok {
			handler(w, req, httprouter.Params{})
		} else if r.NotFound != nil {
			r.NotFound(w, req)
		} else {
			w.Write([]byte(fmt.Sprintf("not found:%v", path)))
		}
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func NewRouter() *Router {
	return &Router{
		middlewareChain: make([]middleware, 0),
		mux:             make(map[string]httprouter.Handle),
	}
}

func (r *Router) Use(m middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}

func (r *Router) Add(path string, handle httprouter.Handle) {
	var mergeHandler = handle
	/*for _, m := range r.middlewareChain {
		//函数嵌套的部分其实是在这里实现的
		mergeHandler = m(mergeHandler)
	}*/
	//上面的代码是不能使用的，因为是函数嵌套，所以最后加入的会被先执行
	for i := len(r.middlewareChain) - 1; i > 0; i-- {
		mergeHandler = r.middlewareChain[i](mergeHandler)
	}
	r.mux[path] = mergeHandler
}

func Phello(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("hello[%s]", params.ByName("name"))))
}

//计算耗时
func PStatisticsTimeMiddle(next httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		start := time.Now()
		next(w, req, params)
		timeSlapsed := time.Since(start)
		log.Printf("用户[%v]%s-请求[%s],耗时:%v\n", params.ByName("name"), req.Method, req.URL.Path, timeSlapsed)
	})
}

//计算请求书
func PCountReq(next httprouter.Handle) httprouter.Handle {
	counts := make(map[string]uint)
	return httprouter.Handle(func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		path := req.URL.Path
		next(w, req, params)
		log.Printf("用户[%v]%s-请求[%s],次数:%v\n", params.ByName("name"), req.Method, path, counts[path])
	})
}

func main() {
	port := ":9002"
	r := NewRouter()
	r.Use(PCountReq)
	r.Use(PStatisticsTimeMiddle)
	r.Add("/", Phello)
	http.ListenAndServe(port, r)
}
