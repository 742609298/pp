package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	bir := "1992-10-29"
	log.Println("bir:", bir)

	t := time.Now().Unix()
	starttime := time.Unix(int64(t), 0) //golang强语言类型，这里要求t必须是int64，如非int64需强制转为int64
	startT := starttime.Format("2006-01-02")
	fmt.Println(startT)

	birlist := strings.Split(bir, "-")
	nowlist := strings.Split(startT, "-")
	birint, _ := strconv.Atoi(birlist[0])
	nowint, _ := strconv.Atoi(nowlist[0])
	log.Println(birint)
	log.Println(nowint)

}
