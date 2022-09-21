package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	reallyDone = 1
	doneEating = 0
)

// Msg holds a value that allows the host and the philosophers to communicate
type Msg int

// ChopS is a chopstick
type ChopS struct{ sync.Mutex }

// Philo is a philosopher
type Philo struct{ leftCS, rightCS *ChopS }

// AnswerChannel is for communication between host and philosopher
type AnswerChannel chan bool

// RunHost runs the host that communicates with the philosophers through channels.
func RunHost(requests chan AnswerChannel, done chan int, numPhilosophers int, wg *sync.WaitGroup) {
	counterDone := 0
	numEating := 0

	for {
		select {
		case ch := <-requests:
			ch <- numEating < 2
		case y := <-done:
			if y == doneEating {
				numEating--
			} else if y == reallyDone {
				counterDone++
				if counterDone >= numPhilosophers {
					wg.Done()
					return
				}
			}
		}
	}
}

func (p Philo) eat(i int, requestChannel chan AnswerChannel, done chan int) {

	ch := make(AnswerChannel)

	for n := 0; n < 3; {
		requestChannel <- ch
		if <-ch {
			n++
			choice := rand.Intn(2) // 0 or 1
			if choice == 0 {
				p.leftCS.Lock()
				p.rightCS.Lock()
			} else {
				p.rightCS.Lock()
				p.leftCS.Lock()
			}

			fmt.Println("starting to eat ", i)
			fmt.Println("finishing eating ", i)

			p.leftCS.Unlock()
			p.rightCS.Unlock()
			done <- doneEating
		}
	}

	done <- reallyDone
}

func main() {
	rand.Seed(42)

	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
	}

	requests := make(chan AnswerChannel)
	doneChannel := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1) // The host will run until all philosophers are done eating
	go RunHost(requests, doneChannel, 5, &wg)

	for i := 0; i < 5; i++ {
		go philos[i].eat(i, requests, doneChannel)
	}
	wg.Wait()
}
