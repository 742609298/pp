package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main()  {
	log.Print("post请求--begin")
	resp, err := http.PostForm("https://api.megvii.com/faceid/liveness/v2/get_token",
		url.Values{
			"appId": {"iWY3Afa6faXmxuPgXYiQzgfI3BiROh70"},
			"appKey": {"m3qYWTh6OunkdaaPfHDu9yHnrtt1Bjfj"},
			"name":{"https://www.baidu.com/"},
			"idNum":{"https://www.baidu.com/"},
			"cardNo":{"123"},
			"mobile":{"何绪君"},
		})
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
	
}