package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/polds/imgbase64"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)
const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

var coder = base64.NewEncoding(base64Table)
func main(){



	img, _ := imgbase64.FromLocal("C:\\Users\\soft\\Desktop\\he.jpg")

	log.Println("imag:",img)



	log.Print("post请求--begin")
	resp, err := http.PostForm("https://api-cn.faceplusplus.com/cardpp/v1/ocridcard",
		url.Values{"api_key": {"bfaiC9IEvyrwMmVaAlFeQuLkNBjMgcUF"},
			"api_secret": {"x6Ou7wF3W4HxL5JMJMwqpgobVvmB9YlN"},
			"image_base64":{img},
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

	reskyc1 := make(map[string]interface{})
	_ = json.Unmarshal(body, &reskyc1)

	cards := reskyc1["cards"]
	list := cards.([]interface{})
	mapres :=list[0].(map[string]interface{})
	name := mapres["name"].(string)
	idcardnum := mapres["id_card_number"].(string)
	log.Println("name:",name)
	log.Println("id:",idcardnum)

}
