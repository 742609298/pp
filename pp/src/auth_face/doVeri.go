package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Print("get请求--begin")

	resp, err := http.Get("https://api.megvii.com/faceid/lite/do?token=abd2e7538a31ab0a7615d55f423579f7")
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
