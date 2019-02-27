package coinsms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"unsafe"
)

//account: I4301462
//password: FG5PHAri8v780e
//url: http://intapi.253.com/send/json
//'发送国际短信  格式 ： 86手机号
func SendSMS_GJ(phone string, msg string) (string, error) {
	params := make(map[string]interface{})
	// API账号，50位以内。必填
	params["account"] = "I4301462"
	// API账号对应密钥，联系客服获取。必填
	params["password"] = "FG5PHAri8v780e"
	// 手机号码，格式(区号+手机号码)，例如：8615800000000，其中86为中国的区号
	params["mobile"] = phone
	//短信内容。长度不能超过536个字符。必填
	params["msg"] = msg //"【253云通讯】您好，您的验证码是999999"
	//用户收到短信之后显示的发件人，国内不支持自定义，国外支持，但是需要提前和运营商沟通注册，具体请与TIG对接人员确定。选填
	params["senderId"] = ""
	bytesData, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err.Error())
		return "fail", err
	}
	reader := bytes.NewReader(bytesData)
	url := "http://intapi.253.com/send/json"
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
