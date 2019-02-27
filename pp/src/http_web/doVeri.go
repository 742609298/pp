package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Print("get请求--begin")

	resp, err := http.Get("https://api.megvii.com/faceid/lite/do?token=bd88f128a7289c36d853ce324d846dc5")
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
