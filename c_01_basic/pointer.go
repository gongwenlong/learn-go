package main

import "fmt"

// Go 语言为程序员提供了控制数据结构的指针的能力；但是，你不能进行指针运算。
func main()  {
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int
	intP = &i1
	//符号 * 可以放在一个指针前，如 *intP，那么它将得到这个指针指向地址上所存储的值；这被称为反引用（或者内容或者间接引用）操作符；另一种说法是指针转移。
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)


	s := "good bye"
	var p *string = &s
	*p = "ciao"
	// s p 指向同一块区域
	fmt.Printf("Here is the pointer p: %p\n", p) // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s) // prints same string

}
