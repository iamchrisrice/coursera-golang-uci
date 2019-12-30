package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numberOfConcurrentEaters = 2

type host struct {
	eaters chan bool
}

const numberOfTimesEachPhilosopherEats = 3

type philosopher struct {
	id int
}

type chopstick struct {
}

var chopsticks = sync.Pool{
	New: func() interface{} {
		return new(chopstick)
	},
}

func (p philosopher) eat(h *host) {
	defer wg.Done()

	for i := 0; i < numberOfTimesEachPhilosopherEats; i++ {
		<-h.eaters

		fmt.Println("starting to eat", p.id)

		leftChopstick := chopsticks.Get()
		rightChopsticks := chopsticks.Get()

		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		chopsticks.Put(leftChopstick)
		chopsticks.Put(rightChopsticks)

		fmt.Println("finished eating", p.id)

		h.eaters <- true
	}
}

var wg sync.WaitGroup

func (h host) begin() {
	for i := 0; i < numberOfConcurrentEaters; i++ {
		h.eaters <- true
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	host := new(host)
	host.eaters = make(chan bool, numberOfConcurrentEaters)

	go host.begin()

	philosophers := make([]*philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &philosopher{i}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philosophers[i].eat(host)
	}

	wg.Wait()
}
