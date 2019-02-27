package main

import (
	"fmt"
	"github.com/menduo/gobaidumap"
)

func main(){
	var lat string = "19.9139658858"
	var lng string = "110.2107238770"
	//从坐标到地址
	locationToaddress, err := gobaidumap.GetAddressViaGEO(lat, lng)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("坐标地址：", locationToaddress)
		fmt.Println("坐标地址-地址", locationToaddress.Result.AddressComponent)
	}
	//有地址到坐标
	address := "广州市天河区石牌小学"
	addressToLocation, err := gobaidumap.GetGeoViaAddress(address)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("从地址到坐标-all", addressToLocation)
		fmt.Println("从地址到坐标 - Lat", addressToLocation.Result.Location.Lat)
		fmt.Println("从地址到坐标 - Lng", addressToLocation.Result.Location.Lng)
		fmt.Println("\n")
	}
	//由ip到地址
	//ipAddress := "222.76.214.60"
	ipAddress := "218.77.129.195"
	IPToAddress, err := gobaidumap.GetAddressViaIP(ipAddress)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("从ip到地址：", IPToAddress)
		fmt.Println("从ip到地址-地址", IPToAddress.Content.Address)
		fmt.Println("\n")
	}
}
