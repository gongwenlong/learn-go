package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	// go getData(ch)
	fmt.Println(<-ch) // prints only 0

	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}

/**
package main

import "fmt"

func main() {
    ch := make(chan string)
    go sendData(ch)
    getData(ch)
}

func sendData(ch chan string) {
    ch <- "Washington"
    ch <- "Tripoli"
    ch <- "London"
    ch <- "Beijing"
    ch <- "Tokio"
    close(ch)
}

func getData(ch chan string) {
    for {
        input, open := <-ch
        if !open {
            break
        }
        fmt.Printf("%s ", input)
    }
}

 */

