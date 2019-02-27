package main

import (
	"fmt"
	"github.com/menduo/gobaidumap"
)

func main(){
	//ipAddress :="103.88.46.222"
	//ipAddress :="223.223.195.196"
	//ipAddress := "218.77.129.195"
	ipAddress := "127.0.0.1"
	IPToAddress, err := gobaidumap.GetAddressViaIP(ipAddress)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("从ip到地址：", IPToAddress)
		fmt.Println("从ip到地址-地址", IPToAddress.Content.Address)
		fmt.Println("\n")
	}


}
