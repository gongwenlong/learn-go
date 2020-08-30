package main

import "fmt"

var num int = 10
var numx2, numx3 int

func main() {
	println("In main before calling greeting")
	greeting()
	println("In main after calling greeting")

	// 多返回值
	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()


	var min, max int
	min, max = MinMax(78, 65)
	fmt.Printf("Minmium is: %d, Maximum is: %d\n", min, max)

	// 传递变长参数
	x := min1(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	slice := []int{7,9,3,5,1}
	x = min1(slice...)
	fmt.Printf("The minimum in the slice is: %d", x)

}

func greeting() {
	println("In greeting: Hi!!!!!")
}


func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3
	return
}

func MinMax(a int, b int) (min int, max int) {
	if a < b {
		min = a
		max = b
	} else { // a = b or a > b
		min = b
		max = a
	}
	return
}

func min1(s ...int) int {
	if len(s)==0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

// 计算函数执行时间
// start := time.Now()
//longCalculation()
//end := time.Now()
//delta := end.Sub(start)
//fmt.Printf("longCalculation took this amount of time: %s\n", delta)