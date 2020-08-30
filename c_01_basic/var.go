package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
)

// Go和许多编程语言不同
// 它在声明变量时将变量的类型放在变量的名称之后
// 它是为了避免像 C 语言中那样含糊不清的声明形式，
// 例如：int* a, b;。在这个例子中，
// 只有 a 是指针而 b 不是。
// 如果你想要这两个变量都是指针，则需要将它们分开书写
// 而在 Go 中，则可以很轻松地将它们都声明为指针类型：var a, b *int


//当一个变量被声明之后，系统自动赋予它该类型的零值：
// int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil。
// 所有的内存在 Go 中都是经过初始化的


// 可见性
// 变量的命名规则遵循骆驼命名法，即首个单词小写，每个新单词的首字母大写
// 例如：numShips 和 startDate。
// 但如果你的全局变量希望能够被外部包所使用，则需要将首个单词的首字母也大写

var Pi1 float64

func main()  {
	var a int = 15
	var i = 5
	var b bool = false
	var str string = "Go says hello to the world!"

	var aa = 15
	var bb = false
	var str1 = "Go says hello to the world!"

	/*
		var (
		a = 15
		b = false
		str = "Go says hello to the world!"
		numShips = 50
		city string
	)
	 */

	fmt.Println(a)
	fmt.Println(i)
	fmt.Println(b)
	fmt.Println(str)

	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(str1)

	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)

	// 值类型和引用类型
	// int、float、bool 和 string 这些基本类型都属于值类型，
	// 数组 和 结构 这些复合类型也是值类型。

	// 指针 属于引用类型， slices ，maps 和 channel 。
	//  被引用的变量会存储在堆中，以便进行垃圾回收，且比栈拥有更大的内存空间。

	fmt.Println("--- 值类型和引用类型 ---")
	fmt.Println(a)
	fmt.Println(&a)
	j := i
	fmt.Println(j)
	fmt.Println(&j)

	// 如果你声明了一个局部变量却没有在相同的代码块中使用它，
	// 同样会得到编译错误

	/**
	  var a, b, c int
	  a, b, c = 5, 7, "abc"
	  a, b, c := 5, 7, "abc"
	  _, b = 5, 7
	 */

	 fmt.Print(Pi1)
}

// init
func init() {
	Pi1 = 4 * math.Atan(1) // init() function computes Pi
}