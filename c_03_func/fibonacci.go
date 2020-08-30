package main

import "fmt"

func main() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

// 优化方式 计算过得存下来
// 或者动态规划

// package main
//
//import (
//    "fmt"
//    "time"
//)
//
//const LIM = 41
//
//var fibs [LIM]uint64
//
//func main() {
//    var result uint64 = 0
//    start := time.Now()
//    for i := 0; i < LIM; i++ {
//        result = fibonacci(i)
//        fmt.Printf("fibonacci(%d) is: %d\n", i, result)
//    }
//    end := time.Now()
//    delta := end.Sub(start)
//    fmt.Printf("longCalculation took this amount of time: %s\n", delta)
//}
//func fibonacci(n int) (res uint64) {
//    // memoization: check if fibonacci(n) is already known in array:
//    if fibs[n] != 0 {
//        res = fibs[n]
//        return
//    }
//    if n <= 1 {
//        res = 1
//    } else {
//        res = fibonacci(n-1) + fibonacci(n-2)
//    }
//    fibs[n] = res
//    return
//}
//
