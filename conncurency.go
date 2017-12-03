package main

import(
	"fmt"
)

/**
 * Example of the Chinese whispering game
 * with 100000 gophers(go routines) 
 * https://youtu.be/f6kdp27TYZs?t=27m24s
 */

func f(left, right chan int) {
	// data flows in directions of arrows
	left <- 1 + <- right
}

func main() {
	const n = 100000
	// create channel of type int
	leftmost := make(chan int)
	right := leftmost
	left  := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}
