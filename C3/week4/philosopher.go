package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Chops struct{ sync.Mutex }

type Philo struct {
	number, fill    int
	leftCS, rightCS *Chops
}

var philos []*Philo

func (p *Philo) eat(wg *sync.WaitGroup) {

	p.leftCS.Lock()
	p.rightCS.Lock()
	fmt.Println("Starting to eat ", p.number)
	p.fill++
	p.leftCS.Unlock()
	p.rightCS.Unlock()
	fmt.Println("Finishing eating ", p.number)

	wg.Done()
}

func host() {
	for {
		curr := 0
		var arr []int
		total := 0
		for i := 0; i < 5; i++ {
			total += philos[i].fill
		}
		if total == 15 {
			break
		}
		for i := 0; i < 5; i++ {
			if curr >= 2 {
				break
			}
			if philos[i].fill < 3 {
				arr = append(arr, i)
				curr++
			}
		}

		wg.Add(curr)
		for _, num := range arr {

			go philos[num].eat(&wg)
		}
		wg.Wait()
	}
}

func main() {
	CSticks := make([]*Chops, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(Chops)
	}

	philos = make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i + 1, 0, CSticks[i], CSticks[(i+1)%5]}
	}
	host()

}
