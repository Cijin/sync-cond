package main

import (
	"fmt"
	"sync"
	"time"
)

type Donation struct {
	cond    *sync.Cond
	balance int
}

var donation = &Donation{
	cond: sync.NewCond(&sync.Mutex{}),
}

var f = func(goal int) {
	donation.cond.L.Lock()
	defer donation.cond.L.Unlock()

	for donation.balance < goal {
		donation.cond.Wait()
	}

	fmt.Println("Goal reached:", goal)
}

func main() {
	go f(5)
	go f(10)

	for {

		time.Sleep(time.Second)
		donation.cond.L.Lock()
		donation.balance++
		donation.cond.L.Unlock()
		donation.cond.Broadcast()
	}
}
