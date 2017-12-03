package main

import (
	"fmt"
	"time"
)
/**
 * A map, that would reseumble a list of redis cache keys
 */
var cache = map[string]bool {
    "r1": true,
    "r3": true,
    "r4": true,
}
/**
 * [fanIn, uses Select(similar to switch) in order to pick the input func that finishes first)]
 * @return {chan} []
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
 * [fromCache look in cache, for request path]
 * @return
 */
func fromCache(r request) <-chan string {
	c := make(chan string)
	/**
	 * @todo, confused why this dosen't always win
	 */
	go func() {
		if cache[r.path] && ! r.auth {
	  	c <- fmt.Sprintf("Winner: 🔥 Cache")
		}
	}()
	return c
}
/**
 * [fromStore, nothing to see here just a random return time]
 * @return
 */
func fromStore(request) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Winner: 📁 store")
			time.Sleep(time.Duration(5000))
		}
	}()
	return c
}
/**
 * request
 */
type request struct {
    path string
    auth bool
}
/**
 * [Make 5 fictious requests and see which comes back cache or store]
 */
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Loop #", i)
		r := fmt.Sprintf("r%v", i)
		c := fanIn(
			fromCache(request{r, false}), 
			fromStore(request{r, false}),
		)
		fmt.Println(<-c)
		fmt.Println("===")
	}

	fmt.Println(`All done`)
}
