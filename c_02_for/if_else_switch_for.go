package main

import "fmt"

func main() {
	var first int = 10
	var cond int

	/**
	if else
	 */
	if first <= 0 {

		fmt.Printf("first is less than or equal to 0\n")
	} else if first > 0 && first < 5 {

		fmt.Printf("first is between 0 and 5\n")
	} else {

		fmt.Printf("first is 5 or greater\n")
	}
	if cond = 5; cond > 10 {

		fmt.Printf("cond is greater than 10\n")
	} else {

		fmt.Printf("cond is not greater than 10\n")
	}


	/**
	switch
	 */

	var num1 int = 100

	switch num1 {
	case 98, 99:
		fmt.Println("It's equal to 98")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98 or 100")
	}

	/**
	for
	 */

	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for ix :=0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	}
	str2 := "日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix :=0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}

	/**
	break
	 */
	var i int = 5
	for {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
		if i < 0 {
			break
		}
	}

	/**
	continue
	 */
	for j := 0; j < 10; j++ {
		if j == 5 {
			continue
		}
		fmt.Println(j)
		fmt.Print(" ")
	}

	/**
	LABEL
	func main() {
    i:=0
    HERE:
        print(i)
        i++
        if i==5 {
            return
        }
        goto HERE
}
	 */
}
