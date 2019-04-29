package cyclicbarrier

import (
	"sync"
	"testing"
	"time"
)

func checkBarrier(t *testing.T, cb *CyclicBarrier,
	expectedParties, expectedWaiting int, expectedIsBroken bool) {

	parties, numberWaiting := cb.GetParties(), cb.GetWaiting()
	isBroken := cb.IsBroken()

	if expectedParties >= 0 && parties != expectedParties {
		t.Error("barrier must have parties = ", expectedParties, ", but has ", parties)
	}
	if expectedWaiting >= 0 && numberWaiting != expectedWaiting {
		t.Error("barrier must have numberWaiting = ", expectedWaiting, ", but has ", numberWaiting)
	}
	if isBroken != expectedIsBroken {
		t.Error("barrier must have isBroken = ", expectedIsBroken, ", but has ", isBroken)
	}
}

// All goroutines wait for cyclic barries
func TestAwaitDone(t *testing.T) {
	n := 100 // goroutines count
	cb := New(n)
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			cb.Await()
			wg.Done()
		}()
	}

	wg.Wait()
	checkBarrier(t, cb, n, 0, false)
}

// Break cyclic barries
func TestAwaitBreak(t *testing.T) {
	n := 100 // goroutines count
	cb := New(n)
	cb2 := New(n)
	for i := 0; i < n-1; i++ {
		go func() {
			cb.Await()
			cb2.Await()
		}()
	}

	time.Sleep(1 * time.Second)
	cb.BreakBarrier()
	cb2.Await()
	checkBarrier(t, cb, n, 1, true)
}
