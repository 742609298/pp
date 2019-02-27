package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

//定义发送邮件内容的结构体
type SenEmailMsg struct {
	from     string
	pwd      string
	to       string
	nickname string
	title    string
	msg      string
}

//发送邮件
func main() {
	//_w.Header().Set("Access-Control-Allow-Origin", "*") //允许跨域
	//定义返回结构体
	type Res struct {
		Code int      `json:"code"`
		Msg  string   `json:"msg"`
		Data []string `json:"data"`
	}
	//result := &Res{}

	//var para map[string]string         //map结构体
	//body:=make(map[string]interface{})//ioutil.ReadAll(_r.Body) //接收参数
	//log.Print("register-emailcode-requestparams-", string(body[:]))
	//json.Unmarshal(body, &para) //解析参数并存至para
	//log.Print("register-emailcode-requestbody-", para)
	//接收参数
	sendEmailMsg := &SenEmailMsg{}
	sendEmailMsg.from = "3404925259@qq.com"
	sendEmailMsg.pwd = "smhgsvvvhxdwcije"
	sendEmailMsg.to = "742609298@qq.com"
	sendEmailMsg.nickname = "lovingcoinex"
	sendEmailMsg.title = "lovingcoinex邮箱注册"
	sendEmailMsg.msg = `
		<html>
		<body  style="width:582;margin:0;padding:0">
		<div>
			<div style="width:582px;height:52px;background-color:#14476b">
				<h3 style="margin:10px;font-size:30px;color:#fff";text-align:center;line-height:80px;">LOVING-COINEX邮箱注册</h3>
			</div>
			<div style="padding:0 20px">
				
			<p style="font-size:14px">尊敬的用户您好</p>
			<p>点击以下链接来验证您的Email</p>
			<a href="http://user.lovingcoinex.net/#/user/register?token=` + sendEmailMsg.to + `&type=1">http://user.lovingcoinex.net/#/user/register?token=` + sendEmailMsg.to + `&type=1</a>
			<p style="width:520px;">
				如果单击该链接没有反应，您可以将该链接复制粘贴到您浏览器的地址栏中，或在地址栏中重新键入该链接。您返回user.lovingcoinex.net之后。我们将为您提供进一步的说明。
			</p>
			</div>
		</div>
		</body>
		</html>
		`
	//验证邮箱格式是否合法

	//检查邮箱是否已注册
	//userCheckEmail := &userdb.User{}




		//SendMsg(sendEmailMsg.from, sendEmailMsg.pwd, sendEmailMsg.to, sendEmailMsg.nickname, sendEmailMsg.title, sendEmailMsg.msg, "html")
		SendMsg(sendEmailMsg.from, sendEmailMsg.pwd, sendEmailMsg.to, sendEmailMsg.nickname, sendEmailMsg.title, sendEmailMsg.msg, "html")
		//result.Code = 200
		//result.Msg = "success"
		//log.Print("register-emailcode-resulr-", result.Msg)
		//data, _ := json.Marshal(result)
		//fmt.Fprintf(_w, string(data))


}

//发送邮件 __formEmailAdd 发件人邮箱 _password 发件人授权码 toEmailAdd 对方邮箱地址 ,_nickName 发件人名称,_title 邮件标题,_msg 邮件内容
func SendMsg(_fromEmailAdd string, _password string, _toEmailAdd string, _nickName string, _title string, _msg string, mailtype string) {
	fmt.Println("===================方法接收参数====================")
	fmt.Println(_fromEmailAdd, _password, _toEmailAdd, _nickName, _title, _msg)
	fmt.Println("===================方法接收参数====================")
	// 邮箱地址
	UserEmail := _fromEmailAdd
	// 端口号，:25也行
	Mail_Smtp_Port := ":587"
	//邮箱的授权码，去邮箱自己获取
	Mail_Password := _password
	// 此处填写SMTP服务器
	Mail_Smtp_Host := "smtp.qq.com"
	auth := smtp.PlainAuth("", UserEmail, Mail_Password, Mail_Smtp_Host)
	to := []string{_toEmailAdd}
	nickname := _nickName
	user := UserEmail

	subject := _title
	content_type := "Content-Type: text/" + mailtype + "; charset=UTF-8"
	body := _msg
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail(Mail_Smtp_Host+Mail_Smtp_Port, auth, user, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	}
}
