package main

import (
	"fmt"
	"time"
)

func main() {
	// make an Add2 function, give it a name p2, and call it:
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	// make a special Adder function, a gets value 2:
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))


	var f = Adder1()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Println(f(300))

	var g int
	go func(i int) {
		s := 0
		for j := 0; j < i; j++ { s += j }
		g = s
	}(1000)

	time.Sleep(1e9)
	fmt.Println(g)

}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func Adder1() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
