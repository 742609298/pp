package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)
var counter1 int = 0
func main() {
	log.Print("协程begin。。。。")
	lock := &sync.Mutex{}

	for i:=0; i<10; i++ {
		go Count1(lock,i)
	}

	time.Sleep(5*time.Second)
}
func Add1(x, y int) {
	z := x + y
	fmt.Println(z)
}
func Count1(lock *sync.Mutex,index int) {
	//打印线程的id

	//lock.Lock()
	counter1++
	for i:=0; ; i++ {
		fmt.Println("编号：",index ,"打印", i)
	}
	//fmt.Println("counter1 =", counter1)
	//lock.Unlock()
}
