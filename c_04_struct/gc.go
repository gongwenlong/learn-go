package main

import (
	"fmt"
	"runtime"
)

//Go 开发者不需要写代码来释放程序中不再使用的变量和结构占用的内存，在 Go 运行时中有一个独立的进程，即垃圾收集器（GC），
// 会处理这些事情，它搜索不再使用的变量然后释放它们的内存。可以通过 runtime 包访问 GC 进程。
//
//通过调用 runtime.GC() 函数可以显式的触发 GC，但这只在某些罕见的场景下才有用，
// 比如当内存资源不足时调用 runtime.GC()，它会在此函数执行的点上立即释放一大片内存，
// 此时程序可能会有短时的性能下降（因为 GC 进程在执行）
//

func main()  {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc / 1024)
}

// runtime.SetFinalizer(obj, func(obj *typeObj))