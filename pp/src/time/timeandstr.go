package time

import (
	"fmt"
	"time"
)

func main(){
	//字符串转时间
	d := "2018-10-10 00:00:00"
	stamp, _ := time.ParseInLocation("2006-01-02 15:04:05",d , time.Local)    //使用parseInLocation将字符串格式化返回本地时区时间
	fmt.Println(stamp.Unix())   //再转为Unix时间戳

	//字符串转时间
	d2 := "2018-10-10 23:59:59"
	//s2,_:= time.Parse("2006-01-02 03:04:05",d2)
	s2, _ := time.ParseInLocation("2006-01-02 15:04:05",d2 , time.Local)
	fmt.Println(s2.Unix())   //再转为Unix时间戳

	//时间转字符串
	t := 1527134541
	starttime := time.Unix(int64(t), 0)  //golang强语言类型，这里要求t必须是int64，如非int64需强制转为int64
	startT := starttime.Format("2006-01-02 03:04:05")
	fmt.Println(startT)

}