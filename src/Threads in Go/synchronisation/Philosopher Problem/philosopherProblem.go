package main

import (
	"fmt"
	"sync"
)

/********************************************************************************************************
This is the naive/incorrect way of implementing philosophers problem. In this scenario there is a
possibilty of deadlock when all the philosophers/goroutines pick up the chopstick to their left.
********************************************************************************************************/

var wg sync.WaitGroup

type chopS struct {
	sync.Mutex
}

type philo struct {
	leftCS, rightCS *chopS
}

func (p philo) eat() {
	for {
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Println("eating")
		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
	// wg.Done()
}

func main() {
	wg.Add(5)
	cSticks := make([]*chopS, 5)
	for i := 0; i < 5; i++ {
		cSticks[i] = new(chopS)
	}
	philos := make([]*philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &philo{cSticks[i], cSticks[(i+1)%5]}
	}
	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
	wg.Wait()
}
