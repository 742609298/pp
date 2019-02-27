package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//随机生成验证码的方法
func main(){
	//这种方法每次生成的都一样over
	//fmt.Println(rand.Intn(9), ",")
	//fmt.Println(rand.Intn(100))

	//要让伪随机数生成器有确定性，可以给它一个明确的种子。
	//获取时间戳
	timestamp :=time.Now().Unix()
	s1 := rand.NewSource(timestamp)
	r1 := rand.New(s1)
	//调用上面返回的 rand.Source 的函数和调用 rand 包中函数是相同的。
	str :=""
	for i :=0;i<=5;i++{
		str = str+ strconv.Itoa( r1.Intn(9))
	}
	fmt.Println(str)

}
