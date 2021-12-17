package main

import (
	"fmt"
	"time"
)

func main() {
	stats := make(map[string]int)

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	for i := 0; i < 20; i++ {
		go func(i int) {
			time.Sleep(10 * time.Second)
			c1 <- fmt.Sprintf("Hello from customer service #1 from iteration #%d", i)
		}(i)

		go func(i int) {
			time.Sleep(8 * time.Second)
			c2 <- fmt.Sprintf("Hello from customer service #2 from iteration #%d", i)
		}(i)

		go func(i int) {
			time.Sleep(6 * time.Second)
			c3 <- fmt.Sprintf("Hello from customer service #3 from iteration #%d", i)
		}(i)

		select {
		case msg1 := <-c1:
			stats["Nick"]++
			time.Sleep(time.Second)
			fmt.Println(msg1)
		case msg2 := <-c2:
			stats["Tim"]++
			time.Sleep(time.Second)
			fmt.Println(msg2)
		case msg3 := <-c3:
			stats["John"]++
			time.Sleep(time.Second)
			fmt.Println(msg3)
		default:
			stats["Customer Waiting"]++
			time.Sleep(2 * time.Second)
			fmt.Println("No customer service is available at this time!")
		}
	}

	fmt.Printf("\n***Customer Service Stats***\n%v", stats)

	close(c1)
	close(c2)
	close(c3)
}
