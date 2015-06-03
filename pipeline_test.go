package txt2web

import (
	"fmt"
	"time"
)

// based upon http://blog.golang.org/pipelines

// next is we need a process with two sub nodes, such that additional data is passed
// first phase returns two channels, second phase expects two channels and returns one as usual

func Example_pipeline() {
	term(sq(sum(gen(1, 2, 3, 4, 5))))

	// output:
	// 9
	// 49
	// 25
}

// generator
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// wait is intended to pass one int per second.
// however, the overall effect is that term prints all values after 4 seconds
// ok, that's because fmt.Println buffers, so, wait works perfectly fine
func wait(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n
			time.Sleep(time.Second)
		}
		close(out)
	}()
	return out
}

// sum reduces the number of ints in the channel by summing every two values
// interestingly enough, this node works better with odd values in the channel
// an even number adds a trailing zero (unless we check for oka)
func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for {
			a, oka := <-in
			b, okb := <-in
			// with even numbers in the chan, oka will be false first
			if oka {
				out <- a + b
			}
			if !okb {
				break
			}
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// terminal node
func term(in <-chan int) {
	// no need for a goroutine here as as we want the term on the stack
	// as long as there are ints in the channel.
	for n := range in {
		// use println in case there is a 'wait' in the chain
		// println(n)
		fmt.Println(n)
	}
}
