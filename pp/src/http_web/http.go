package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main(){
	log.Print("post请求--begin")
	resp, err := http.PostForm("https://api.megvii.com/faceid/liveness/v2/get_token",
		url.Values{"api_key": {"iWY3Afa6faXmxuPgXYiQzgfI3BiROh70"},
			"api_secret": {"m3qYWTh6OunkdaaPfHDu9yHnrtt1Bjfj"},
			"return_url":{"https://www.baidu.com/"},
			"notify_url":{"https://www.baidu.com/"},
			"biz_no":{"123"},
			"idcard_name":{"何绪君"},
			"idcard_number":{"370112199210294517"},
			"comparison_type":{"1"},
			//"idcard_mode":{"2"},
			//"":{""},
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
