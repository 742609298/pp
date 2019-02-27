package main

import (
	"github.com/shopspring/decimal"
	"log"
	"xunhuan1"
)

func main()  {
	log.Print(xunhuan1.A)
	//
	//a :=decimal.New(1.090918,10)
	//b:=decimal.New(0.03340102,10)
	c,_ :=decimal.NewFromString("1.090918")
	d,_:=decimal.NewFromString("0.03340102")
	//log.Println(a)
	//log.Println(b)
	log.Println(c)
	log.Println(d)
	//log.Println(a.Mul(b))
	log.Println(c.Mul(d))
}
