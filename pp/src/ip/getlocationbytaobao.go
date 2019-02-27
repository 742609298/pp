package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	resp, err := http.Get("http://ip.taobao.com/service/getIpInfo.php?ip="+"192.168.46.222")//103.88.46.222
	log.Println(err)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
	location  :=make(map[string]interface{})
	_ = json.Unmarshal(body, &location)
	log.Println(location)
	log.Println(location["data"])
	log.Println(location["data"].(map[string]interface{})["country"])
	log.Println(location["data"].(map[string]interface{})["city"])
}
