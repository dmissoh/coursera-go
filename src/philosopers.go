package main

import (
	"fmt"
	"sync"
	"time"
)

var waitingGroup sync.WaitGroup

var host = make(chan bool, 2)

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	id             int
	leftChopStick  *ChopStick
	rightChopStick *ChopStick
}

func (philosopher Philosopher) eat() {

	// a philosopher eats 3 time
	for i := 0; i < 3; i++ {

		host <- true

		philosopher.leftChopStick.Lock()
		philosopher.rightChopStick.Lock()

		fmt.Println("start eating: ", philosopher.id+1)
		time.Sleep(1 * time.Second)
		fmt.Println("stop eating: ", philosopher.id+1)

		philosopher.leftChopStick.Unlock()
		philosopher.rightChopStick.Unlock()

		<-host
	}

	// done eating
	waitingGroup.Done()
}

func main() {

	// create 5 chop sticks
	chopSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = new(ChopStick)
	}

	// create 5 philosophers
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i, chopSticks[i], chopSticks[(i+1)%5]}
	}

	// start eating
	for i := 0; i < 5; i++ {
		waitingGroup.Add(1)
		go philosophers[i].eat()
	}

	// wait until all philosopher are done eating
	waitingGroup.Wait()
}
