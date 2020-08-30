package main

import "fmt"

// 数组是具有相同 唯一类型 的一组已编号且长度固定的数据项序列（这是一种同构的数据结构）；
// 这种类型可以是任意的原始类型例如整型、字符串或者自定义类型。
// 数组长度必须是一个常量表达式，并且必须是一个非负整数。数组长度也是数组类型的一部分，
// 所以 [5] int 和 [10] int 是属于不同类型的。数组的编译时值初始化是按照数组顺序完成的（如下）。
//
// 注意事项 如果我们想让数组元素类型为任意类型的话可以使用空接口作为类型

// Go 语言中的数组是一种 值类型（不像 C/C++ 中是指向首元素的指针），
// 所以可以通过 new() 来创建： var arr1 = new([5]int)。
// 那么这种方式和 var arr2 [5]int 的区别是什么呢？
// arr1 的类型是 *[5]int，而 arr2 的类型是 [5]int
//

func main() {
	var arr1 [5]int
	arr1p := &arr1

	for i:=0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}
	fmt.Println(arr1)
	fmt.Println(&arr1)
	fmt.Println(arr1p)
	fmt.Println(&arr1p)

	for i:=0; i < len(arr1); i++ {
		fmt.Printf("Array at index %d is %d\n", i, arr1[i])
	}

	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

	var ar [3]int
	f(ar)   // passes a copy of ar
	fp(&ar) // passes a pointer to ar


	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	// var arrAge = [5]int{18, 20, 15, 22, 16}
	// var arrLazy = [...]int{5, 6, 7, 8, 22}

	// var arrKeyValue = []string{3: "Chris", 4: "Ron"}

	for i:=0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}

	for i := 0; i < 3; i++ {
		fp(&[3]int{i, i * i, i * i * i})
	}

	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array) // Note the explicit address-of operator
	// to pass a pointer to the array
	fmt.Printf("The sum of the array is: %f", x)

}

func f(a [3]int) { fmt.Println(a) }
func fp(a *[3]int) { fmt.Println(a) }

func Sum(a *[3]float64) (sum float64) {
	for _, v := range a { // derefencing *a to get back to the array is not necessary!
		sum += v
	}
	return
}
