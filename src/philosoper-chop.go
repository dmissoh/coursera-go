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

func (p Philosopher) eat() {

	for i := 0; i < 3; i++ {

		host <- true

		p.leftChopStick.Lock()
		p.rightChopStick.Lock()

		fmt.Println("starting to eat: ", p.id+1)
		time.Sleep(2 * time.Second)
		fmt.Println("finishing eating: ", p.id+1)

		p.leftChopStick.Unlock()
		p.rightChopStick.Unlock()

		<-host

	}
	waitingGroup.Done()
}

func main() {

	chopSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = new(ChopStick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i, chopSticks[i], chopSticks[(i+1)%5]}
	}

	for i := 0; i < 5; i++ {
		waitingGroup.Add(1)
		go philosophers[i].eat()
	}

	waitingGroup.Wait()

}
