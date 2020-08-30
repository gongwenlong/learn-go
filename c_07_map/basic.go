package main

import "fmt"

// map 是一种特殊的数据结构：一种元素对（pair）的无序集合，
// pair 的一个元素是 key，对应的另一个元素是 value，
// 所以这个结构也称为关联数组或字典。
// 这是一种快速寻找值的理想结构：给定 key，对应的 value 可以迅速定位。
//
// map 这种数据结构在其他编程语言中也称为字典（Python）、hash 和 HashTable 等。
// map 是引用类型

func main() {
	var mapLit map[string]int
	//var mapCreated map[string]float32
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapAssigned["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])

	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Println(mf)

}
