package main

import "fmt"

func main() {
	ch := make(chan int)
	defer fmt.Println("defer: ", <-ch)
	ch <- 42
	fmt.Println("main: ", <-ch)
}
