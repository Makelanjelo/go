package main

import (
	"fmt"
	// "time"
)

func main() {

	// Асинхроннность / рутины
	ch := make(chan int)

	go say("Hello Go", ch)
	// time.Sleep(2 * time.Second)
	fmt.Println(<-ch)

}

func say(greet string, ch chan int) {
	fmt.Println(greet)
	ch <- 7
}
