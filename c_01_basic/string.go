package main

import (
	"fmt"
	"strings" // strings 和 strconv
)

// 解释字符串：
// 该类字符串使用双引号括起来，其中的相关的转义字符将被替换，这些转义字符包括：
/*
  \n：换行符
  \r：回车符
  \t：tab 键
  \u 或 \U：Unicode 字符
  \\：反斜杠自身
 */

// 非解释字符串：该类字符串使用反引号括起来，支持换行，例如：
/*
`This is a raw string \n` 中的 `\n\` 会被原样输出
 */

// 字符串的内容（纯字节）可以通过标准索引法来获取，在中括号 [] 内写入索引，索引从 0 开始计数：
//
//字符串 str 的第 1 个字节：str[0]
//第 i 个字节：str[i - 1]
//最后 1 个字节：str[len(str)-1]

func main()  {
	str1 := "Beginning of the string " +
		"second part of the string"

	fmt.Println(str1)

	s := "hel" + "lo,"
	s += "world!"
	fmt.Println(s) //输出 “hello, world!”

	// 在循环中使用加号 + 拼接字符串并不是最高效的做法，更好的办法是使用函数
	// strings.Join()
	// 有没有更好地办法了？有！使用字节缓冲（bytes.Buffer）
	// 拼接更加给力

	var str string = "Hi, I'm Marc, Hi."

	fmt.Printf("The position of \"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Marc"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))

}