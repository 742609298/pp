package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Print("get请求--begin")

	resp, err := http.Get("https://api.megvii.com/faceid/lite/get_result?api_key=iWY3Afa6faXmxuPgXYiQzgfI3BiROh70" +
		"&api_secret=m3qYWTh6OunkdaaPfHDu9yHnrtt1Bjfj" +
		"&biz_id=1546828149,7d17c962-68c5-40d9-9ee8-02467d90b7ac" +
		"&return_image=0")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
	//fmt.Println(decoder.ConvertString(string(body)))


	//log.Print(string(body))
	//log.Print(body)
	log.Print("\u4f55\u7eea\u541b")

}
