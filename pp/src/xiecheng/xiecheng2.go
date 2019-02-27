package main

import (
	"log"
)
var num int
func Count2(ch chan int,i int) {
	//num++
	//time.Sleep(5*time.Second)
	log.Print("Counting:::::::",i,"::::",num)//不一定打印10次为什么
	ch <- i
}

func main() {

	chs := make([] chan int, 10)

	for i:=0; i<10; i++ {
		chs[i] = make(chan int)//无缓存的channel会进行阻塞
		go Count2(chs[i],i)
	}

	for _, ch := range(chs) {
		x:= <- ch
		log.Print("result:",x)//执行10次
	}
	//time.Sleep(5)



}

//func main() {
//	a:=1
//	b:=2
//	log.Println("a",a)
//	log.Println("b",b)
//}

