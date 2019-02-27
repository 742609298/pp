package main

import (
	"log"
)

func main()  {
	log.Println("begin:")
	defer func() {
		if err := recover(); err != nil {
			log.Print("捕获异常:", err)
			// 存在异常回滚
		}else{
			//没有异常commit
			log.Print("没有异常:", err)
		}
		log.Print("main--over")
	}()


	F()
	log.Println("F()--over")

	F2()
	//panic("a")//后面的defer不执行
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3)
	//defer fmt.Println(4)

}

func F() {
	//自己捕获异常
	defer func() {
		if err := recover(); err != nil {
			log.Print("捕获异常:", err)
			// 存在异常回滚
		}else{
			//没有异常commit
		}

	}()
	//抛出异常
	panic("f")



}

func F2(){
	panic("f2")
}


