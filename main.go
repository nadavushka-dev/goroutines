package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	msg string
}

func main() {
	start := time.Now()

	respch := make(chan *Order, 3)
	wg := sync.WaitGroup{}

	wg.Add(3)

	go makeBurger(respch, &wg)
	go makeChips(respch, &wg)
	go makeDrink(respch, &wg)

	wg.Wait()
	close(respch)

	for o := range respch {
		fmt.Println(o.msg)
	}
	fmt.Println("Order took: ", time.Since(start))
}

func makeBurger(respch chan *Order, wg *sync.WaitGroup) {
	const makeTime = 300 * time.Millisecond
	time.Sleep(makeTime)

	msg := "Done making burgers"

	respch <- &Order{msg: msg}

	wg.Done()
}

func makeChips(respch chan *Order, wg *sync.WaitGroup) {
	const makeTime = 150 * time.Millisecond
	time.Sleep(makeTime)

	msg := "Done making chips"

	respch <- &Order{msg: msg}

	wg.Done()
}

func makeDrink(respch chan *Order, wg *sync.WaitGroup) {
	const makeTime = 50 * time.Millisecond
	time.Sleep(makeTime)

	msg := "Done making drink"

	respch <- &Order{msg: msg}

	wg.Done()
}
