package syncutil

import (
	"sync"
)

// Waiter is a sync.WaitGroup
// but it also has a Channel() method for use with select{}
type Waiter struct {
	sync.WaitGroup
}

// Waiter.Channel() returns a channel that can be used with select{}
// the channel is closed when WaitGroup.Wait() is satisfied
//
// ex:
//		var w syncutil.Waiter
//		...
//		wch := w.Channel()
//		select {
//			case <-wch: // waiter done
//				...
//			case <-input:
//			...
//		}
func (w *Waiter) Channel() (ch chan int) {
	ch = make(chan int)
	go func() {
		w.Wait()  // wait for WaitGroup
		close(ch) // close channel (NOTE: doesn't happen if Wait() panics)
	}()
	return
}
