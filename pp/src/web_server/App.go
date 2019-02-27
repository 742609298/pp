package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"log"
	"net/http"
	//"random"
	"redisutil"
	"strings"
)
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Wrold!") //这个写入到w的是输出到客户端的
}
func sayhelloName2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Wrold2!") //这个写入到w的是输出到客户端的
}
func getIdentifyingCode(_w http.ResponseWriter, _r *http.Request) {

	code:= random.GetIdentifyingCode()
	cli := redisutil.GetRedisPool()
	conn := cli.Get()
	defer conn.Close()
	//存值
	_, err := conn.Do("SET", "mycode", code, "EX", "60")
	if err != nil {
		_ = json.NewEncoder(_w).Encode("获取验证码失败")
		return
	}

	//fmt.Fprintf(w, "Hello Wrold2!") //这个写入到w的是输出到客户端的
	_ = json.NewEncoder(_w).Encode("【爱币网】您的验证码是："+code)
}
func equalIdentifyingCode(_w http.ResponseWriter, _r *http.Request) {
	type Requestbody struct{

		Code int  `json:"code"`
	}
	body, _ := ioutil.ReadAll(_r.Body)
	requestBody :=Requestbody{}
	cli := redisutil.GetRedisPool()
	conn := cli.Get()
	defer conn.Close()
	if err := json.Unmarshal(body, &requestBody); err != nil {
		log.Print("ERROR: ", err, " - ", string(body[:]))
	}
	log.Print("requestBody", "-", requestBody)
	log.Println("得到的验证码是：",requestBody.Code)
	//取值
	mykeyvalue, err := redis.String(conn.Do("GET", "mycode"))
	log.Println("redis的验证码是：",mykeyvalue)
	if err != nil {
		log.Print("redis get failed:", err)
		_ = json.NewEncoder(_w).Encode("验证码输入错误")
	} else {
		log.Print(mykeyvalue)
		_ = json.NewEncoder(_w).Encode("验证码输入正确")
		return
	}

}

func main() {
	log.Println("========================服务启动==========================")

	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.HandleFunc("/2", sayhelloName2) //设置访问的路由
	http.HandleFunc("/getIdentifyingCode", getIdentifyingCode) //获取验证码
	http.HandleFunc("/equalIdentifyingCode", equalIdentifyingCode) //提交验证码
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Println("========================服务停止==========================")
}
