package main

// import (
// 	// "fmt"
// 	// "time"
// )

func main() {

	// Асинхроннность / рутины

	// ch := make(chan string)
	// ch1 := make(chan int)

	// go say("Hello Go", ch, ch1)
	// // time.Sleep(2 * time.Second)
	// data := <-ch
	// fmt.Println(data, <-ch1)

	// ch := make(chan int)

	// go say("Hello Go", ch)

	// for a := range ch {
	// 	fmt.Println(a)
	// }
	// fmt.Println(<-ch)
}

// func say(greet string, ch chan string, ch1 chan int) {
// 	fmt.Println(greet)
// 	ch <- "Hello"
// 	ch1 <- 77

// }

// func say(greet string, ch chan int) {
// 	for i := 0; i <= 5; i++ {
// 		ch <- i
// 	}

// 	close(ch) // Закрытие канала

// }
