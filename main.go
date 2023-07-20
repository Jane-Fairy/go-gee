package main

import (
	_ "Adula_Go/contextTest"
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	ch2 := make(chan int, 1)
	go func(chan int) {
		for {
			select {
			case ch <- 0:
				//println("0")
			case ch2 <- 1:
				//println("1")
			}
		}
	}(ch)

	for {
		fmt.Printf("ch %d ", <-ch)
		fmt.Println()
		fmt.Printf("ch %d ", <-ch2)
		fmt.Println()
		//fmt.Println(<-ch2)
	}

}
