package test

import (
	"coinswitch/coinutils/src/github.com/dgrijalva/jwt-go"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type res struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}

//路由中间件，用来验证Token信息
func ValidateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(_w http.ResponseWriter, _r *http.Request) {
		_w.Header().Set("Access-Control-Allow-Origin", "*") //允许跨域
		_w.Header().Add("Access-Control-Allow-Headers", "Authorization")
		_w.Header().Set("content-type", "application/json")
		//获取请求url
		str := _r.RequestURI
		//过滤注册、登陆、钱包调用请求
		if strings.Index(str, "register") > 0 || strings.Index(str, "login") > 0 || strings.Index(str, "bcoinorder") > 0 {
			next.ServeHTTP(_w, _r)
		} else {
			//获取token
			token := _r.Header.Get("Authorization")
			if token != "" {
				//私钥
				SecretKey := []byte("What is Fuck!")
				//token 过期时间

				var expTime int64 = 10000 * 1000 * 60 * 100000000

				var tokenState int //接收Token状态 0-解析token失败 1-token过期 2-正常
				//解析token
				claims, err := ParseToken(token, SecretKey)
				if nil != err {
					fmt.Println("解析失败了", token)
					tokenState = 0
				} else {
					//获取并解析token中的时间
					oldT, _ := strconv.ParseInt(claims.(jwt.MapClaims)["exp"].(string), 10, 64)
					//获取当前时间
					ct := time.Now().UTC().UnixNano()
					//判断token是否过期
					c := ct - oldT
					if c > expTime {
						tokenState = 1
					} else {
						tokenState = 2
					}
				}
				if tokenState == 2 {
					//解析用户id，并将用户id写入请求头
					userId := claims.(jwt.MapClaims)["userId"].(float64) //解析获取userId

					userIdNew := strconv.FormatFloat(userId, 'f', -1, 64) //将userId转换为字符串
					_r.Header.Set("userId", userIdNew)                    //将userId写入请求头传递
					next.ServeHTTP(_w, _r)
				} else if tokenState == 0 {
					res := &res{}
					res.Code = 1
					res.Msg = "未登录"

					ret, _ := json.Marshal(res)

					_, _ = fmt.Fprintf(_w, string(ret))

				} else {
					res := &res{}
					res.Code = 2
					res.Msg = "token过期，请重新登陆"
					ret, _ := json.Marshal(res)

					_, _ = fmt.Fprintf(_w, string(ret))

				}
			}
		}
	})
}

//解析token _tokenSrt：token字符串， _SecretKey密钥
func ParseToken(_tokenSrt string, _SecretKey []byte) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(_tokenSrt, func(*jwt.Token) (interface{}, error) {
		return _SecretKey, nil
	})
	claims = token.Claims

	a := 1.2
	strconv.Itoa(a)
	return
}
