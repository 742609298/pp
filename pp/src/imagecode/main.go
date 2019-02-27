package main

import (
	"encoding/json"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"io/ioutil"
	"log"
	"net/http"
	//"utils"
)

//ConfigJsonBody json request body.
type ConfigJsonBody struct {
	Id          string `json:"id"`
	CaptchaType string `json:"captchaType"`
	VerifyValue string `json:"verifyValue"`
	ConfigAudio     base64Captcha.ConfigAudio //音频
	ConfigCharacter base64Captcha.ConfigCharacter //字母
	ConfigDigit     base64Captcha.ConfigDigit //数字
}

// base64Captcha create http handler
//生成图片验证码
func generateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	//parse request parameters
	//接收客户端发送来的请求参数
	log.Println("生成验证码begin====")

	body, _ := ioutil.ReadAll(r.Body)
	log.Println("生成验证码参数====",string(body[:]))
	requestBody := ConfigJsonBody{}
	if err := json.Unmarshal(body, &requestBody); err != nil {

		log.Println("生成验证码====解析参数出错")
		log.Println(err)
	}

	log.Println("生成验证码====解析参数rb:",requestBody)
	defer r.Body.Close()


	//create base64 encoding captcha
	//创建base64图像验证码

	var config interface{}
	switch requestBody.CaptchaType {
	case "audio":
		//config = postParameters.ConfigAudio
	case "character":
	//config = postParameters.ConfigCharacter
	case "digit":
		config = base64Captcha.ConfigDigit{
			100,
			200,
			6,
			0.5,
			2,
		}

	default:
		//config = postParameters.ConfigDigit
		config = base64Captcha.ConfigDigit{
			10,
			10,
			4,
			0.5,
			2,
		}
	}
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	captchaId, digitCap := base64Captcha.GenerateCaptcha(requestBody.Id, config)
	base64Png := base64Captcha.CaptchaWriteToBase64Encoding(digitCap)

	//or you can do this
	//你也可以是用默认参数 生成图像验证码
	//base64Png := captcha.GenerateCaptchaPngBase64StringDefault(captchaId)

	//set json response
	//设置json响应

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	bodyreturn := map[string]interface{}{"code": 1, "data": base64Png, "captchaId": captchaId, "msg": "success"}
	json.NewEncoder(w).Encode(bodyreturn)
}

// base64Captcha verify http handler
//对比图片验证码
func captchaVerifyHandle(w http.ResponseWriter, r *http.Request) {

	//parse request parameters
	//接收客户端发送来的请求参数
	body, _ := ioutil.ReadAll(r.Body)
	log.Println("生成验证码参数====",string(body[:]))
	requestBody := ConfigJsonBody{}
	if err := json.Unmarshal(body, &requestBody); err != nil {

		log.Println("生成验证码====解析参数出错")
		log.Println(err)
	}

	log.Println("生成验证码====解析参数rb:",requestBody)



	defer r.Body.Close()
	//verify the captcha
	//比较图像验证码
	verifyResult := base64Captcha.VerifyCaptcha(requestBody.Id, requestBody.VerifyValue)

	//set json response
	//设置json响应
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body123 := map[string]interface{}{"code": "error", "data": "验证失败", "msg": "captcha failed"}
	if verifyResult {
		body123 = map[string]interface{}{"code": "success", "data": "验证通过", "msg": "captcha verified"}
	}
	json.NewEncoder(w).Encode(body123)
}

//start a net/http server
//启动golang net/http 服务器
func main() {

	//serve Vuejs+ElementUI+Axios Web Application
	//http.Handle("/", http.FileServer(http.Dir("./static")))

	//api for create captcha
	//创建图像验证码api
	//http.HandleFunc("/api/getCaptcha", generateCaptchaHandler)
	http.HandleFunc("/api/getCaptcha", generateCaptchaHandler) //获取验证码

	//http.HandleFunc("/getIdentifyingCode", getIdentifyingCode) //获取验证码
	//api for verify captcha
	http.HandleFunc("/api/verifyCaptcha", captchaVerifyHandle)//对比图片验证码

	fmt.Println("Server is at localhost:8888")
	//if err := http.ListenAndServe("localhost:8888", nil); err != nil {
	//	log.Fatal(err)
	//}
	err := http.ListenAndServe(":8888", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Println("========================服务停止==========================")
}

//func getIdentifyingCode(_w http.ResponseWriter, _r *http.Request) {
//
//	code := random.GetIdentifyingCode()
//	cli := random.GetRedisPool()
//	conn := cli.Get()
//	defer conn.Close()
//	//存值
//	_, err := conn.Do("SET", "mycode", code, "EX", "60")
//	if err != nil {
//		_ = json.NewEncoder(_w).Encode("获取验证码失败")
//		return
//	}
//
//	//fmt.Fprintf(w, "Hello Wrold2!") //这个写入到w的是输出到客户端的
//	_ = json.NewEncoder(_w).Encode("【爱币网】您的验证码是：" + code)
//}
