package main

import (
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup
var wg2 sync.WaitGroup
var ch = make(chan int, 2)

type cStick struct {
	sync.Mutex
}

type philo struct {
	lCStick, rCStick *cStick
	id               int
}

func (p philo) eat() {
	for i := 0; i < 3; i++ {
		p.lCStick.Lock()
		p.rCStick.Lock()
		fmt.Println("starting to eat", p.id)
		fmt.Println("finishing eating", p.id)
		time.Sleep(10 * time.Millisecond)
		p.rCStick.Unlock()
		p.lCStick.Unlock()
	}
	ch <- 1
	wg2.Done()
}

func host(philos []*philo) {
	wg2.Add(5)
	// Channels have been used to make sure only two philosophers eat at one time
	go philos[0].eat()
	go philos[2].eat()
	<-ch
	<-ch
	go philos[1].eat()
	go philos[3].eat()
	<-ch
	<-ch
	go philos[4].eat()

	wg2.Wait()
	wg1.Done()
}

func main() {
	wg1.Add(1)
	cSticks := make([]*cStick, 5)
	for i := 0; i < 5; i++ {
		cSticks[i] = new(cStick)
	}
	philos := make([]*philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &philo{cSticks[i], cSticks[(i+1)%5], i + 1}
	}
	go host(philos)
	wg1.Wait()
}
