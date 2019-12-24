net/http不能满足么，为什么要用这个扩展？
    
* 不能单独的对请求方法(GET,POST,PUT,DELETE)注册特定的处理函数
* 不支持path变量参数，通配符
* 不能自动对path进行校准
* 扩展性不足，比如group路由，v1,v2这种
    

1 如果在两个路由中拥有一致的http方法和请求路径前缀，且在某个位置出现了A路由是wildcard(通配符)参数，B路由是普通字符串，那么就会产生路由冲突

    GET /user/info/:id
    GET /user/:id
    
2  httprouter考虑字典树的深度，所以在路由中的参数不能超过255，否则会导致后面的参数会无法解析

3  httprputer还支持 \* 号来进行通配，不过 \* 号开头的参数只能放在路由的结尾

    /src/*css
    
4  定制404
    
    //404捕获(zhyi这里的捕获不包括静态资源的)
   	r.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
   		w.Write([]byte("not found:" + r.URL.String()))
   	})
    
5  httprouter的handler是自定义的，并不是http.Handler，httprouter的handler多个一个参数解析：
    
    type Handle func(http.ResponseWriter, *http.Request, Params)
    
   但两者是兼容的，可以平滑升级，无需担心
   
6  访问静态资源：

    r := httprouter.New()
    r.ServeFiles("/static/*filepath", http.Dir("./static"))

7 定制panic

    //异常捕获,返回500错误
    r.PanicHandler = func(w http.ResponseWriter, r *http.Request, v interface{}) {
        //这里有输出的话，响应状态就是200了
        //w.Write([]byte(fmt.Sprintf("500 error:%s", v)))
        w.WriteHeader(http.StatusInternalServerError)
    }
    
原理介绍
   
httprouter和众多衍生的router使用的数据结构被称为压缩字典树(Radix Tree),普通的字典树对每个字母都需要建立
一个孩子节点，这样会导致字典树的层数比较深。压缩字典树每个节点上不只存储一个字母，而是一个path字符串，可以减少
树的层数


net/http的Handlr、HandlerFunc、ServerHTTP到底是什么关系，好混乱,他们是什么关系:


    type Handler interface {
        ServerHTTP(ResponseWriter, *Request)
    }
    
    type HandlerFunc func(ResponseWriter, *Request)
    
    func (f HandlerFunc) ServerHTTP(w ResponseWriter, r *Request) {
        f(w, r)
    }

为什么要使用中间件

为了将业务代码和非业务代码进行分离，一般的非业务需求都是在http请求处理前和处理后做一些事情
中间件通过一个或者多个函数包装handler，再返回一个包括了各个中间件逻辑的新的handler的函数
链

*[非优雅的写法](https://github.com/lltty/note/tree/master/go/http_middle.go)
*[优雅写法](https://github.com/lltty/note/tree/master/go/http_middle_perfect.go)


