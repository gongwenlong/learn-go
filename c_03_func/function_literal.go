package main

import "fmt"

// 当我们不希望给函数起名字的时候，
// 可以使用匿名函数，例如：func(x, y int) int { return x + y }。
//
// 这样的一个函数不能够独立存在（编译器会返回错误：non-declaration statement outside function body），
// 但可以被赋值于某个变量，即保存函数的地址到变量中：fplus := func(x, y int) int { return x + y }，
// 然后通过变量名对函数进行调用：fplus(3,4)。
//
// 当然，您也可以直接对匿名函数进行调用：func(x, y int) int { return x + y } (3, 4)
//
func main() {
	f0()
	fmt.Println(f1())
}

func f0() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) } //此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}

//关键字 defer 经常配合匿名函数使用，它可以用于改变函数的命名返回值。
//
//匿名函数还可以配合 go 关键字来作为 goroutine 使用
func f1() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

// 当您在分析和调试复杂的程序时，无数个函数在不同的代码文件中相互调用，
// 如果这时候能够准确地知道哪个文件中的具体哪个函数正在执行，
// 对于调试是十分有帮助的。您可以使用 runtime 或 log 包中的特殊函数来实现这样的功能。
// 包 runtime 中的函数 Caller() 提供了相应的信息，
// 因此可以在需要的时候实现一个 where() 闭包函数来打印函数执行的位置

