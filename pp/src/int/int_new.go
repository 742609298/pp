package main

import (
	"fmt"
	"log"
	"strconv"
)

func main(){
	// int 转 string
	i := int64(-2048)
	fmt.Println(strconv.FormatInt(i, 2))  // -100000000000
	fmt.Println(strconv.FormatInt(i, 8))  // -4000
	fmt.Println(strconv.FormatInt(i, 10)) // -2048
	fmt.Println(strconv.FormatInt(i, 16)) // -800
	fmt.Println(strconv.FormatInt(i, 36)) // -1kw
	// int 转 string
	strconv.Itoa(int(i))//默认按照10进制
	fmt.Println("++++++++++++++++++++++=") // -1kw

	var a uint  = 3
	var b string = ""
	b = strconv.Itoa(int(a))
	log.Print(b)


}
