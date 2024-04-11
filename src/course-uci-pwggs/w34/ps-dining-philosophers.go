// Course 3 week 3
package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	NumMeals        = 3
	NumPhilosophers = 5
	maxEating       = 2
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	idx     int
	meals   int
	leftCS  *ChopS
	rightCS *ChopS
}

func (p Philo) eat(c chan *Philo, wg *sync.WaitGroup) {
	for i := 0; i < NumMeals; i++ {
		c <- &p
		if p.meals < NumMeals {
			p.leftCS.Lock()
			p.rightCS.Lock()
			fmt.Println("starting to eat", p.idx)
			p.meals++ // eating
			fmt.Println("finishing eating", p.idx)
			p.rightCS.Unlock()
			p.leftCS.Unlock()
			wg.Done()
		}
	}
}

func main() {

	fmt.Printf("[%v philosophers, each will eat %v times, max concurrency is %v]\n", NumPhilosophers, NumMeals, maxEating)
	chSticks := PlaceChopsticks()
	philosophers := PlacePhilosophers(chSticks)

	var wg sync.WaitGroup
	wg.Add(NumPhilosophers * NumMeals)

	c := make(chan *Philo, maxEating)
	go Host(c)

	for i := 0; i < NumPhilosophers; i++ {
		go philosophers[i].eat(c, &wg)
	}

	wg.Wait()
	fmt.Printf("[Completed!]")
}

// PlaceChopsticks -
func PlaceChopsticks() []*ChopS {
	cs := make([]*ChopS, NumPhilosophers)
	for i := 0; i < NumPhilosophers; i++ {
		cs[i] = new(ChopS)
	}
	return cs
}

// PlacePhilosophers -
func PlacePhilosophers(cs []*ChopS) []*Philo {
	p := make([]*Philo, NumPhilosophers)
	for i := 0; i < NumPhilosophers; i++ {
		p[i] = &Philo{i + 1, 0, cs[i], cs[(i+1)%NumPhilosophers]}
	}
	return p
}

// Host -
func Host(c chan *Philo) {
	for {
		if len(c) == maxEating {
			<-c
			<-c
			time.Sleep(50 * time.Millisecond)
		}
	}
}
