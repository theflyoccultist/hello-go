// Philosopher's problem, exploring concurrency in Go

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numPhilosophers = 5
	numForks        = 5
	numMeals        = 3
)

type Philosopher struct {
	id        int
	leftFork  *sync.Mutex
	rightFork *sync.Mutex
	ladle     *sync.Mutex
}

func (p *Philosopher) eat(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numMeals; i++ {
		// think
		think := rand.Intn(5) + 1
		fmt.Printf("Philosopher %d is thinking for %d seconds\n", p.id, think)
		time.Sleep(time.Duration(think) * time.Second)

		// pick up ladle
		p.ladle.Lock()
		fmt.Printf("Philosopher %d used the ladle\n", p.id)

		// pick up fork
		p.leftFork.Lock()
		fmt.Printf("Philosopher %d picked up left fork\n", p.id)
		p.rightFork.Lock()
		fmt.Printf("Philosopher %d picked up right fork\n", p.id)

		// eat after picking up two forks
		eat := rand.Intn(5) + 1
		fmt.Printf("Philosopher %d is eating for %d seconds\n", p.id, eat)
		time.Sleep(time.Duration(eat) * time.Second)

		// put down forks
		p.leftFork.Unlock()
		fmt.Printf("Philosopher %d put down the left fork\n", p.id)
		p.rightFork.Unlock()
		fmt.Printf("Philosopher %d put down the right fork\n", p.id)

		// Put down ladle
		p.ladle.Unlock()
		fmt.Printf("Philosopher %d put down the ladle\n", p.id)
	}
}

func main() {
	forks := make([]*sync.Mutex, numForks)
	for i := range forks {
		forks[i] = &sync.Mutex{}
	}

	ladle := &sync.Mutex{}

	philosophers := make([]*Philosopher, numPhilosophers)
	for i := range philosophers {
		leftFork := forks[i]
		rightFork := forks[(i+1)%numForks]
		philosophers[i] = &Philosopher{id: i + 1, leftFork: leftFork, rightFork: rightFork, ladle: ladle}
	}

	var wg sync.WaitGroup
	wg.Add(numPhilosophers)
	for _, p := range philosophers {
		go p.eat(&wg)
	}
	wg.Wait()
}
