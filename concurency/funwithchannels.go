//passes messages between 2 goroutines.  I wanted to test to see if I could load a struct with a bunch of buffered channels
//so  that it could be passed to a function. 

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type chans struct {
	ch1 chan string
	ch2 chan string
	ch3 chan string
	ch4 chan int
}

func main() {

	ch1 := make(chan string, 100)
	ch2 := make(chan string, 100)
	ch3 := make(chan string, 100)
	ch4 := make(chan int, 100)

	ch := chans{}
	ch.ch1 = ch1
	ch.ch2 = ch2
	ch.ch3 = ch3
	ch.ch4 = ch4
	for i := 0; i < 100; i++ {
		ch.ch4 <- (rand.Intn(3) + 1)
	}

	stringpush(ch)

	go func() {
		for {

			select {
			case x := <-ch.ch1:
				fmt.Println(x)
				ch.ch4 <- rand.Intn(3) + 1
			case x := <-ch.ch2:
				fmt.Println(x)
				ch.ch4 <- rand.Intn(3) + 1
			case x := <-ch.ch3:
				fmt.Println(x)
				ch.ch4 <- rand.Intn(3) + 1

			}

		}
	}()

	for {

	}

}

func stringpush(ch chans) {

	go func() {
		for {
			holder := <-ch.ch4
			if 1 == holder {
				time.Sleep(250 * time.Millisecond)
				ch.ch1 <- "channel 1 - Howdy"

			} else if 2 == holder {
				time.Sleep(256 * time.Millisecond)
				ch.ch2 <- "channel 2 - Hi!"
			} else if 3 == holder {
				time.Sleep(244 * time.Millisecond)
				ch.ch3 <- "channel 3 - Hello?"
			}

		}
	}()

}
