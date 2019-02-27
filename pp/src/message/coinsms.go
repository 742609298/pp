package coinsms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"unsafe"
)

/*  发送一个短信内容 */

type JsonPostSample struct {
}

//account: N3004060
//password: sZXk39pxHTa75c
//url: http://smssh1.253.com/msg/send/json

//account: I4301462
//password: FG5PHAri8v780e
//url: http://intapi.253.com/send/json
//发送国内短信
func SendSMS(phone string, msg string) (string, error) {
	params := make(map[string]interface{})

	//请登录zz.253.com获取API账号、密码以及短信发送的URL
	//params["account"] = "N3004060"  //创蓝API账号
	//params["password"] = "sZXk39pxHTa75c" //创蓝API密码
	params["account"] = "N3004060"        //创蓝API账号
	params["password"] = "sZXk39pxHTa75c" //创蓝API密码
	params["phone"] = phone               //手机号码

	//设置您要发送的内容：其中“【】”中括号为运营商签名符号，多签名内容前置添加提交
	params["msg"] = msg //url.QueryEscape("【253云通讯】您好，您的验证码是999999")
	params["report"] = "true"
	bytesData, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err.Error())
		return "fail", err
	}
	reader := bytes.NewReader(bytesData)
	url := "http://smssh1.253.com/msg/send/json" //短信发送URL guonei

	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return "fail", err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return "fail", err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "fail", err
	}

	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
	return "", nil
}
