package main

import (
	"log"

)
import (
	"gopkg.in/ini.v1"
"./configutil"
)

//var  filepath ="C:\\Users\\soft\\Desktop\\pp\\resource\\test.config"
var  filepath2 ="./resource/test.config" //使用相对路径，从pp路径开始
type Config2 struct {   //配置文件要通过tag来指定配置文件中的名称
	MqttHostname string  `ini:"mqtt_hostname"`
	MqttPort string  `ini:"mqtt_port"`
	MqttUser string  `ini:"mqtt_user"`
	MqttPass string  `ini:"mqtt_pass"`
	MqttKeepalive	int    `ini:"mqtt_keepalive"`
	MqttTimeout  	int    `ini:"mqtt_timeout"`
}

func main(){

	Configutil.Conf2()
	config,err := ReadConfig(filepath)   //也可以通过os.arg或flag从命令行指定配置文件路径
	if err != nil {
		log.Fatal(err)
	}
	log.Println(config)


}
//读取配置文件并转成结构体
func ReadConfig2(path string) (Config, error) {
	var config Config
	conf, err := ini.Load(path)   //加载配置文件
	if err != nil {
		log.Println("load config file fail!")
		return config, err
	}
	conf.BlockMode = false
	err = conf.MapTo(&config)   //解析成结构体
	if err != nil {
		log.Println("mapto config file fail!")
		return config, err
	}
	return config, nil
}
