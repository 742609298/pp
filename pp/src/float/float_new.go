package main

import (
	"log"
	"strconv"
)

func main(){
	//string     float
	stringA :="123.06"
	floatA, err := strconv.ParseFloat(stringA, 64)
	if err != nil {
		panic(err)
	}
	log.Print("floatA:",floatA)

	stringB :="123.06"
	floatB, err := strconv.ParseFloat(stringB, 64)

	if err != nil {
		panic(err)
	}
	log.Print("floatB:",floatB)

	//float 转 string
	//fmt：格式标记（b、e、E、f、g、G）
	//// 格式标记：
	// 'b' (-ddddp±ddd，二进制指数)
	// 'e' (-d.dddde±dd，十进制指数)
	// 'E' (-d.ddddE±dd，十进制指数)
	// 'f' (-ddd.dddd，没有指数)
	// 'g' ('e':大指数，'f':其它情况)
	// 'G' ('E':大指数，'f':其它情况)
	//// 如果格式标记为 'e'，'E'和'f'，则 prec 表示小数点后的数字位数
	// 如果格式标记为 'g'，'G'，则 prec 表示总的数字位数（整数部分+小数部分）
	//prec如果是负数，按实际参数显示
	stringSum := strconv.FormatFloat(floatB, 'f', 3, 64)
	log.Println("stringSum:",stringSum)

}
