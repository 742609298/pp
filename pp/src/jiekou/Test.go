package main

import (
	"log"
)

func main()  {
	var test Tester
	a :=A{name:"x-a"}
	test = &a
	test.speek()
}

type Tester interface {
	speek()
	eat()
}
type A struct {
	name string
}

func (a *A)speek()  {
	log.Println("speek:",a.name)
	log.Println("speek:",a)//&{x-a}
	log.Println("speek:",&a)//0xc000006028
	log.Println("speek:",*a)//{x-a}
}
func (a *A)eat()  {
	log.Println("eat:",a.name)
}


