package main

import (
	"log"
	"github.com/robfig/cron"
	"time"
)

func main() {

	i := 0
	c := cron.New()
	//5秒执行一次
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {//异步执行
		i++
		task(i)
	})
	c.Start()

	select{}
	//defer c.Stop()


}
func task(i int){


	log.Println("cron running:", i)
	time.Sleep(10*time.Second)
	log.Println("cron running:", i*10)
}
