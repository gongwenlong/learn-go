package main

import "fmt"

const Pi = 3.14159

// 显式类型定义
const b string = "abc"

// 隐式类型定义：
//const b = "abc"

/**
常量的值必须是能够在编译时就能够确定的；
你可以在其赋值表达式中涉及计算过程，
但是所有用于计算的值必须在编译期间就能获得
 */
const c1 = 4/3

/*
因为在编译期间自定义函数均属于未知，
因此无法用于常量的赋值，
但内置函数可以使用，如：len ()。

const c2 = getNumber()
func getNumber() int {
	return 1
}
*/

//数字型的常量是没有大小和符号的，
//并且可以使用任何精度而不会导致溢出
const Ln2= 0.693147180559945309417232121458176568075500134360255254120680009
const Log2E= 1/Ln2 // this is a precise reciprocal
const Billion = 1e9 // float constant
const hardEight = (1 << 100) >> 97

// 并行赋值
const beef, two, c = "eat", 2, "veg"
//const Monday, Tuesday, Wednesday, Thursday,
// Friday, Saturday = 1, 2, 3, 4, 5, 6
//const (
//	Monday, Tuesday, Wednesday = 1, 2, 3
//	Thursday, Friday, Saturday = 4, 5, 6
//)

// 枚举
const (
	Unknown = 0
	Female = 1
	Male = 2
)

// 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 aa=0, bb=1, cc=2

const (
	aa = iota
	bb = iota
	cc = iota
)



func main() {
	fmt.Println(Pi)
	fmt.Println(b)
	fmt.Println(c1)
	fmt.Println(Ln2)
	fmt.Println(Log2E)
	fmt.Println(Billion)
	fmt.Println(hardEight)
	fmt.Println(beef)
	fmt.Println(two)
	fmt.Println(c)

	fmt.Println(Unknown)
	fmt.Println(Female)
	fmt.Println(Male)

	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)


}

