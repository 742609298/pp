package main

import (
	"io"
	"log"
	"net"
	"net/http"
)


var a int = 0
var b int = 0
func hello(w http.ResponseWriter, r *http.Request) {
	//测试一下协程
	//测试获取ip
	log.Print("ip:::::",r.RemoteAddr)//192.168.31.168:58008
	addrs,_ := net.InterfaceAddrs()
	log.Print("addrs:",addrs)
	a = a+1
	log.Println("a:",a)
	go test_xiecheng()
	io.WriteString(w, "Hello world!")
}
func test_xiecheng(){
	b = b+1
	log.Println("b:",b)

	//log.Print("go1")
	//time.Sleep(10*time.Second)
	//log.Print("go2")
	//for{
	//	log.Print("go2")
	//}
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	log.Print("============开启服务=================")
	server := http.Server{

		Addr:    ":8889",
		Handler: &myHandler{},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = hello

	server.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("============接受到请求=================")
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "My server: "+r.URL.String())
}
