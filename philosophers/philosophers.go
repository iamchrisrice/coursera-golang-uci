package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numberOfConcurrentEaters = 2
const numberOfPhilosophers = 5
const numberOfTimesEachPhilosopherEats = 3

type host struct {
	eaters chan bool
}

type philosopher struct {
	id int
}

type chopstick struct {
}

var chopsticks = make(chan chopstick, numberOfPhilosophers)

func (p philosopher) eat(h *host) {
	defer wg.Done()

	for i := 0; i < numberOfTimesEachPhilosopherEats; i++ {
		<-h.eaters

		left, right := <-chopsticks, <-chopsticks

		fmt.Println("starting to eat", p.id)

		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		fmt.Println("finished eating", p.id)

		chopsticks <- right
		chopsticks <- left

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

	philosophers := make([]*philosopher, numberOfPhilosophers)
	for i := 0; i < numberOfPhilosophers; i++ {
		philosophers[i] = &philosopher{i + 1}
		chopsticks <- *new(chopstick)
	}

	for i := 0; i < numberOfPhilosophers; i++ {
		wg.Add(1)
		go philosophers[i].eat(host)
	}

	wg.Wait()
}
