package cyclicbarrier

// Package cyclicbarrier provides an implementation of Cyclic Barrier primitive.
import (
	"sync"
)

// CyclicBarrier is a synchronizer that allows a set of goroutines to wait for each other
// to reach a common execution point, also called a barrier.
// CyclicBarriers are useful in programs involving a fixed sized party of goroutines
// that must occasionally wait for each other.
type CyclicBarrier struct {
	lock   sync.RWMutex
	p      int
	n      int
	b      chan struct{}
	broken bool
}

// New initializes a new instance of the CyclicBarrier,
// specifying the number of parties
func New(parties int) *CyclicBarrier {
	return &CyclicBarrier{
		lock:   sync.RWMutex{},
		p:      parties,
		n:      parties,
		b:      make(chan struct{}),
		broken: false,
	}
}

func (cb *CyclicBarrier) reset(broken bool) {
	cb.lock.Lock()
	defer cb.lock.Unlock()

	cb.broken = broken
	close(cb.b)
	cb.b = make(chan struct{})
}

// Init reinit cyclic barrier for future use
func (cb *CyclicBarrier) Init() {
	cb.lock.Lock()
	defer cb.lock.Unlock()

	cb.n = cb.p
	close(cb.b)
	cb.b = make(chan struct{})
}

// Await waits until all parties have invoked await on this barrier.
func (cb *CyclicBarrier) Await() {
	cb.lock.Lock()
	cb.n--
	n := cb.n
	b := cb.b
	cb.lock.Unlock()

	if n > 0 {
		<-b
	} else {
		cb.reset(false)
	}
}

// GetParties returns the itotal number of parties.
func (cb *CyclicBarrier) GetParties() int {
	cb.lock.RLock()
	defer cb.lock.RUnlock()

	return cb.p
}

// GetWaiting returns the number of parties currently waiting at the barrier.
func (cb *CyclicBarrier) GetWaiting() int {
	cb.lock.RLock()
	defer cb.lock.RUnlock()

	return cb.n
}

// IsBroken queries if this barrier is in a broken state.
func (cb *CyclicBarrier) IsBroken() bool {
	cb.lock.RLock()
	defer cb.lock.RUnlock()

	return cb.broken
}

// BreakBarrier set barrier to broken state, all clients Await will be ended
// For broken state check IsBroken() result
func (cb *CyclicBarrier) BreakBarrier() {
	cb.reset(true)
}
