package main

import (
	"fmt"
	"gittest/calc"
)

func main()  {
	res:=calc.Add(10,20)
	fmt.Println(res)
	res=calc.Sub(50,20)
	fmt.Println(res)
}