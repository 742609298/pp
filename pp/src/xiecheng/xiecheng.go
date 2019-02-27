package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
)
var counter int = 0
func main() {
	log.Print("协程begin。。。。")
	lock := &sync.Mutex{}

	for i:=0; i<10; i++ {
		go Count(lock)
	}

	for {
		lock.Lock()

		c := counter

		lock.Unlock()

		runtime.Gosched()// 出让时间片

		if c >= 10 {
			break
		}
	}
}
func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}
func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println("counter =", counter)
	lock.Unlock()
}
