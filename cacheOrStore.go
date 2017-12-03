package main

import (
	"fmt"
	"time"
	"math/rand"
)

/**
 * [main description]
 * @return {[type]} [description]
 */
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
				case s := <-input1: c <- s
				case s := <-input2: c <- s				
			}
		}
	}()
	return c
}
/**
 * [main description]
 * @return {[type]} [description]
 */
func printString(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Winner: %s", msg)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
/**
 * [main description]
 * @return {[type]} [description]
 */
func main() {
	c := fanIn(printString(`fromCache`), printString(`fromStore`))
	for i := 0; i < 10; i++ {
		fmt.Println("Loop #", i)
		fmt.Println(<-c)
		fmt.Println("===")
	}
	fmt.Println(`All done`)
}
