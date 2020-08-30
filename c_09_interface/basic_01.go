package main

import "fmt"

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

/* method to determine the value of a stock position */
func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

/* method to determine the value of a car */
func (c car) getValue() float32 {
	return c.price
}

/* contract that defines different things that have value */
type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

func main() {
	var o valuable = stockPosition{"GOOG", 577.20, 4}
	showValue(o)
	o = car{"BMW", "M3", 66500}
	showValue(o)
}

// 一个标准库的例子
//
// io 包里有一个接口类型 Reader:
//
// type Reader interface {
//    Read(p []byte) (n int, err error)
// }
//    var r io.Reader
//    r = os.Stdin    // see 12.1
//    r = bufio.NewReader(r)
//    r = new(bytes.Buffer)
//    f,_ := os.Open("test.txt")
//    r = bufio.NewReader(f)

// 嵌套
// type ReadWrite interface {
//    Read(b Buffer) bool
//    Write(b Buffer) bool
//}
//
//type Lock interface {
//    Lock()
//    Unlock()
//}
//
//type File interface {
//    ReadWrite
//    Lock
//    Close()
//}
