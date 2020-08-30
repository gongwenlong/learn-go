package main

import "fmt"


type ByteSize float64
const (
	_ = iota // 通过赋值给空白标识符来忽略值
	KB ByteSize = 1<<(10*iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)


type TZ int


func main() {
	/*
	  Go 是强类型语言，因此不会进行隐式转换，任何不同类型之间的转换都必须显式说明
	  var a int
	  var b int32
	  a = 15
	  b = a + a    // 编译错误
	  b = b + 5    // 因为 5 是常量，所以可以通过编译
	 */

	var n int16 = 34
	var m int32
	// compiler error: cannot use n (type int16) as type int32 in assignment
	//m = n
	m = int32(n)

	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("16 bit int is: %d\n", n)

	//Go 对于值之间的比较有非常严格的限制，只有两个类型相同的值才可以进行比较，
	// 如果值的类型是接口（interface），它们也必须都实现了相同的接口。
	// 如果其中一个值是常量，那么另外一个值的类型必须和该常量类型相兼容的。
	// 如果以上条件都不满足，则其中一个值的类型必须在被转换为和另外一个值的类型相同之后才可以进行比较。

	var aVar = 10
	fmt.Println(aVar == 5)
	fmt.Println(aVar == 10)

	// || && | &
	/*
	    优先级     运算符
	    7      ^ !
		6      * / % << >> & &^
		5      + - | ^
		4      == != < <= >= >
		3      <-
		2      &&
		1      ||
	*/

	fmt.Println(GB)


	// 类型别名
	var a, b TZ = 3, 4
	c := a + b
	fmt.Printf("c has the value: %d", c) // 输出：c has the value: 7


	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3) // UTF-8 code point

}