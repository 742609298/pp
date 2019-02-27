package main

import (
	"flag"
	"fmt"
	"os"
)

func main2() {
	//定义Flag
	//方式一：通过flag.String(), Bool(), Int() 等flag.Xxx()方法，该种方式返回一个相应的指针
	namePtr := flag.String("name", "Anson", "user's name")
	agePtr := flag.Int("age", 22, "user's age")
	vipPtr := flag.Bool("vip", true, "is a vip user")
	//方式二：通过flag.XxxVar()方法将flag绑定到一个变量，该种方式返回值类型
	var email string
	flag.StringVar(&email, "email", "abc@gmail.com", "user's email")
	//还有第三种方式，通过flag.Var()绑定自定义类型，自定义类型需要实现Value接口(Receives必须为指针)
	//flag.Var(&flagVal, "name", "help message for flagname")

	//解析命令行参数,值保存到定义的flag
	flag.Parse()

	//调用Parse解析后，就可以直接使用flag本身(指针类型)或者绑定的变量了(值类型)
	//还可通过flag.Args(), flag.Arg(i)来获取非flag命令行参数
	others := flag.Args() //保存Flag以外的变量
	fmt.Println("name:", *namePtr)
	fmt.Println("age:", *agePtr)
	fmt.Println("vip:", *vipPtr)
	fmt.Println("email:", email)
	fmt.Println("other:", others)
	fmt.Println("---------")
	for i := 0; i < len(flag.Args()); i++ {
		fmt.Println("Arg", i, "=", flag.Arg(i))
	}

	// 改变默认的 Usage
	flag.Usage = Usage

}
//帮助文档
func Usage() {
	fmt.Fprintf(os.Stderr, `news-server version: 1.10.0Usage: match-server [-config filename]Options:`)
	flag.PrintDefaults()
}
