package txt2web

import (
	"fmt"
	"time"
)

// based upon http://blog.golang.org/pipelines

// next is we need a process with two sub nodes, such that additional data is passed
// first phase returns two channels, second phase expects two channels and returns one as usual

// dit lijkt een beetje op doorgeven van de quit channel (maar dat is dat weer voor iedere node)

// variadic return ... ? https://code.google.com/p/go/issues/detail?id=119

/*
fan-out, fan in
---------------

Een node kan een fan-out implementeren, dat is, de inkomende stroom
wordt verdeeld over meerdere uitgaande stromen.

Er zijn twee types denkbaar.

In het eerste type bepaalt de node het aantal uitgaande stromen.
Eventueel kan zo'n aantal bepaald worden door een aantal dat als
argument wordt meegegeven. De node retourneert dan de uitgaande stromen.
Dit moet als een slice. Mogelijk bestaat er een variadic return waarde.

In het tweede type krijgt de node de beschikking over uitgaande stromen.
Dat kan niet veel anders dan de uitgaande stromen als (variadic)
argument krijgt. Dan is er vervolgens nog de optie dat deze variadic
argument het is. Maar het is ook denkbaar dat de uitgaande stromen over
de return waarden worden teruggegeven, net als in het eerste type.

> Er is wat discussie rondom variadic return values, maar al met al
> lijkt het er op dat het niet (meer) wordt ondersteund.

Deze discussie is interessant, maar volgens mij niet zo interessant.
*/

func Example_pipeline() {
	term(sq(sum(gen(1, 2, 3, 4, 5))))

	fmt.Println("---")
	fmt.Println("---")

	term(odds_b(odds_a(gen(2, 4, 6))))
	fmt.Println("---")
	term(odds_b(odds_a(gen(1, 4, 6))))
	fmt.Println("---")
	term(odds_b(odds_a(gen(2, 4, 1))))

	fmt.Println("---")
	fmt.Println("---")

	term(odds_c(odds_a(gen(2, 4, 6))))
	fmt.Println("---")
	term(odds_c(odds_a(gen(1, 4, 6))))
	fmt.Println("---")
	term(odds_c(odds_a(gen(2, 4, 1))))

	// `go test -race` mentions race conditions, for now, do not test:
	// o u t p u t:
	// 9
	// 49
	// 25
	// ---
	// ---
	// 2
	// 4
	// channel did not contain odd value
	// 6
	// ---
	// channel contains odd value
	// 1
	// 4
	// 6
	// ---
	// 2
	// 4
	// channel contains odd value
	// 1
	// ---
	// ---
	// 2
	// 4
	// channel did not contain odd value
	// 6
	// ---
	// channel contains odd value
	// 1
	// 4
	// 6
	// ---
	// 2
	// 4
	// channel contains odd value
	// 1
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
			// with even #numbers in the chan, channel will be closed when it is a's
			// turn to receive a value.  So, we have to check oka to prevent a spurious
			// zero in the outbound channel
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

// odds_a is a silly first phase in a two step node.  Function odds_a
// sends true over the bool channel whenever there is an odd value
// in the inbound channel
func odds_a(in <-chan int) (<-chan int, <-chan bool) {
	out := make(chan int)
	hasOdd := make(chan bool)
	go func() {
		noOdds := true
		for n := range in {
			// send exactly one value
			if noOdds && n%2 == 1 {
				noOdds = false
				hasOdd <- true
			}
			out <- n
		}
		close(out)
		// in case you close hasOdd, listeners will receive false values
		// so, apparently no need to send false explicitly
		if noOdds {
			hasOdd <- false
		}
		close(hasOdd)
	}()
	return out, hasOdd
}

// odds_b is a 2nd step function that accepts odds_a as argument.
// it is a little bit undeterministic in that "channel contains odd value"
// come just before the 4.  (in some other attempts, there is no problem)
func odds_b(in <-chan int, hasOdd <-chan bool) <-chan int {
	out := make(chan int)

	// odds_b has two closures, one for each channel.  odds_c uses one
	// closure and a select.

	// combine into one for select??
	// a solution may be in http://stackoverflow.com/questions/13666253
	// but find this a little bit clunky, esp. when I've to check for
	// closed channels

	go func() {
		val := <-hasOdd
		if val {
			fmt.Println("channel contains odd value")
		} else {
			fmt.Println("channel did not contain odd value")
		}
	}()

	go func() {
		for n := range in {
			out <- n
		}
		close(out)
	}()
	return out
}

// odds_c is an alternative for odds_b -  a 2nd step node function
// I expect this version is a little bit more determistic.
func odds_c(in <-chan int, hasOdd <-chan bool) <-chan int {
	out := make(chan int)

	// odds_c uses one closure and a select. odds_b has two closures,
	// one for each channel.

	// see  http://stackoverflow.com/questions/13666253

	go func() {
		for {
			select {
			case odd, odd_open := <-hasOdd:
				if odd_open && odd {
					fmt.Println("channel contains odd value")
				} else if odd_open && !odd {
					fmt.Println("channel did not contain odd value")
				} else {
					hasOdd = nil
				}
			case i, in_open := <-in:
				if in_open {
					out <- i
				} else {
					in = nil
				}
			}

			if hasOdd == nil && in == nil {
				break
			}
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
