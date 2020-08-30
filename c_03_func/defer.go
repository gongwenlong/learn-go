package main

import (
	"fmt"
	"io"
	"log"
)

// 关键字 defer 允许我们进行一些函数执行完成后的收尾工作，例如：
// defer file.Close()

// mu.Lock()
//defer mu.Unlock()

//printHeader()
//defer printFooter()

// defer disconnectFromDB()
func main() {
	function1()
	doDBOperations()
}

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")
	a()
	func1("Go")
	f()
}

func function2() {
	fmt.Printf("function2: Deferred until the end of the calling function!")
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

// 当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）：
func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}


func connectToDB() {
	fmt.Println("ok, connected to db")
}

func disconnectFromDB() {
	fmt.Println("ok, disconnected from db")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB() //function called here with defer
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return //terminate the program
	// deferred function executed here just before actually returning, even if
	// there is a return or abnormal termination before
}

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}
