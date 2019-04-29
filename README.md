# cyclicbarrier

CyclicBarrier is a synchronizer that allows a set of goroutines to wait for each other to reach a common execution point, also called a barrier.

### Usage
Initiate
```go
import "github.com/msaf1980/cyclicbarrier"
...
b1 := cyclicbarrier.New(10) // new cyclic barrier with parties = 10
```
Await
```go
b.Await()    // await other parties
```
Break Barrier
```go
b.BreakBarrier()       // break the barrier
```
Break Barrier
```go
broken := b.IsBroken()       // break barrier status
```

### Simple example
```go
// create a barrier for 10 parties with an action that increments counter
b := cyclicbarrier.New(10)

wg := sync.WaitGroup{}
for i := 0; i < 10; i++ {           // create 10 goroutines (the same count as barrier parties)
    wg.Add(1)
    go func() {
        defer wg.Done

        for j := 0; j < 5; j++ {

            // do some hard work 5 times
            time.Sleep(100 * time.Millisecond)

            b.Await() // ..and wait for other parties on the barrier.
                      // Last arrived goroutine will do the barrier action
                      // and then pass all other goroutines to the next round
        }
    }()
}

wg.Wait()
```
