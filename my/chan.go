package my

import "fmt"

func mains() {
	ch := make(chan int, 1)
	ch <- 1
	<-ch
	fmt.Print("hello world ")
}
