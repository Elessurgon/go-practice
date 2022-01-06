package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	granted_plates          = 3
	numberOfPhilo           = 5
	numberOfConcurrentEater = 2
)

type ChopStick struct {
	id      int
	grabbed sync.Mutex
}

type Philosopher struct {
	idNum  int
	cs     []*ChopStick
	first  int
	second int
}

var (
	// First/Second chopstick (picked up/put down) <-> "left" or "down"
	whichCS = make(map[int]string, 2)
	// Create an ordering for chopsticks
	csticks []*ChopStick
	wg      sync.WaitGroup
	// Buffered channel no more than 2 philosophers can eat concurrently
	hostPermission = make(chan bool, numberOfConcurrentEater)
)

// Randomly determine ordering of chopstick to be picked up/put down for a philosopher
func (p *Philosopher) randCS() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	p.first = r.Intn(2)
	p.second = (p.first + 1) % 2
}

// A randomize timer
func randTimer(t int) {
	timer := rand.Intn(t)
	for {
		if timer <= 0 {
			break
		} else {
			time.Sleep(1 * time.Second)
			timer--
		}

	}
}

// Getting back information "right"/"left" chopstick for a philosopher for
// first/second chopstick to be picked up/put down
func (p *Philosopher) whichCS() {
	if p.first > p.second {
		whichCS[p.first] = "right"
		whichCS[p.second] = "left"
	} else {
		whichCS[p.first] = "left"
		whichCS[p.second] = "right"
	}

}

// Host grant permission with a receive channel (max numberOfConcurrentEater)
func (p *Philosopher) grantedPermision() chan<- bool {
	hostPermission <- true
	fmt.Printf("\n- Host : \"Philosopher <%d> is invited to seat and eat !\"\n", p.idNum)
	return hostPermission
}

// Host revoke permission with a send channel
func (p *Philosopher) revokedPermission() <-chan bool {
	fmt.Printf("\n- Host : \"...Philosopher <%d> is invited to leave his seat !\"\n", p.idNum)
	fmt.Printf("\n- Philosopher <%d> : \"Sure ! I'm leaving the table...\"\n", p.idNum)
	<-hostPermission
	return hostPermission
}

func (p *Philosopher) grabChopsticks() {
	p.randCS()  // generate an ordering
	p.whichCS() // get back info "left"/"right"
	first, second := p.first, p.second
	p.cs[first].grabbed.Lock()
	fmt.Printf("\n...philosopher <%d> picked up %s chopstick %d\n", p.idNum, whichCS[first], p.cs[first].id)
	p.cs[second].grabbed.Lock()
	fmt.Printf("\n...philosopher <%d> picked up %s chopstick %d\n", p.idNum, whichCS[second], p.cs[second].id)
}

func (p *Philosopher) deposeChopsticks() {
	p.whichCS() // get back info "left"/"right"
	first, second := p.first, p.second
	p.cs[first].grabbed.Unlock()
	fmt.Printf("\n...philosopher <%d> put down %s chopstick %d\n", p.idNum, whichCS[first], p.cs[first].id)
	p.cs[second].grabbed.Unlock()
	fmt.Printf("\n...philosopher <%d> put down %s chopstick %d\n", p.idNum, whichCS[second], p.cs[second].id)
}

func initChopSticks() []*ChopStick {
	ChopSticks := make([]*ChopStick, numberOfPhilo)
	for i := 0; i < numberOfPhilo; i++ {
		ChopSticks[i] = new(ChopStick)
		ChopSticks[i].id = i
	}
	return ChopSticks
}

func initHosts() []*Philosopher {
	csticks = initChopSticks()
	philos := make([]*Philosopher, numberOfPhilo)
	for i := 0; i < numberOfPhilo; i++ {
		philos[i] = &Philosopher{
			idNum: i,
			cs: []*ChopStick{
				csticks[i],
				csticks[(i+1)%numberOfPhilo],
			},
			first:  -1,
			second: -1,
		}
	}
	return philos
}

func (p Philosopher) eat() {
	for i := 0; i < granted_plates; i++ {
		p.grantedPermision()
		fmt.Printf("\n...philosopher <%d> is thinking he's hungry !\n", p.idNum)
		p.grabChopsticks()
		fmt.Printf("\n...philosopher <%d> starting to eat its plate NÂ°%d ...\n", p.idNum, i)
		randTimer(3)
		fmt.Printf("\n...philosopher <%d> has finished eating its plate NÂ°%d !\n", p.idNum, i)
		p.deposeChopsticks()

		if i+1 != granted_plates {
			fmt.Printf("\n...philosopher <%d> is thinking : \"Yummy this plate NÂ°%d !...Hoping plate NÂ°%d will be too !\"\n", p.idNum, i, i+1)
			p.revokedPermission()
			randTimer(2)
		} else {
			fmt.Printf("\n...philosopher <%d> is thinking : \"No more plate to eat !\"\n", p.idNum)
			p.revokedPermission()
		}
	}
	wg.Done()
}

func main() {
	fmt.Println("\n- Host : Tonight philosoper's dining, each of you will be invited to taste and eat 3 meals !...")
	fmt.Println("- Host : Please ! Kindly wait to be seated...")
	philos := initHosts()
	for i := 0; i < numberOfPhilo; i++ {
		wg.Add(1)
		go philos[i].eat()
	}
	wg.Wait()
}
