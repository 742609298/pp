package main

import (
	"github.com/shopspring/decimal"
	"log"
)

func main()  {
	//a := decimal.NewFromFloat(3.2)
	b := decimal.NewFromFloat(0)
	//c := a.Div(b)
	//log.Print(c)
	d,_ := decimal.NewFromString("0")
	//e := decimal.NewFromFloat(0)
	log.Println(b==d)
	log.Println(b.IsZero())
	log.Println("over")
	//exp  value后边几个0
	x :=decimal.New(1,10)
	log.Println(x)

	x15 :=decimal.New(1,0)
	decimal15 := x15.Truncate(8).StringFixed(8)
	log.Println(decimal15)



}