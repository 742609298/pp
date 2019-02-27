package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)
//kyc1
func main(){
	log.Print("post请求--begin")
	resp, err := http.PostForm("https://api-cn.faceplusplus.com/cardpp/v1/ocridcard",
		url.Values{"api_key": {"bfaiC9IEvyrwMmVaAlFeQuLkNBjMgcUF"},
			"api_secret": {"x6Ou7wF3W4HxL5JMJMwqpgobVvmB9YlN"},
			"image_url":{"picurl"},
			"legality":{"1"},
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

