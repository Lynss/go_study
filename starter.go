package main

import (
	"./ly"
	"fmt"
)

func main() {
	//a, b, c := ly.AddFirst(1, 2)
	//println(a, b, c)
	//x, y := ly.DirectBack(1, 2)
	//println(x, y)
	//ly.PrintTransformed()
	//ly.ForCommon()
	//ly.ForWhile()
	//ly.ForLoop()
	fmt.Print(ly.Sqrt(9))
	ly.SwitchLike("a")
	ly.SwitchWithoutCondition()

	defer fmt.Println("world")
	fmt.Println("hello")

}
