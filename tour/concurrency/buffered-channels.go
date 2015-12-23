package main

import "fmt"

func move(v int, ch chan int) {
	ch <- v
}

func main() {
	ch := make(chan int, 3)
	move(1, ch)
	move(2, ch)

	x, y := <-ch, <-ch
	fmt.Println(x)
	fmt.Println(y)
}
