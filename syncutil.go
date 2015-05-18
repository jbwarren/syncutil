package syncutil

import (
	"sync"
)

// Waiter is a sync.WaitGroup
// but it also has a WaitChannel() method for use with select{}
type Waiter struct {
	sync.WaitGroup
}

// Waiter.WaitChannel() returns a channel that can be used with select{}
// the channel is closed as soon as WaitGroup.Wait() is satisfied
// so don't call it before you're ready
//
// ex:
//		var w syncutil.Waiter
//		...
//		wch := w.WaitChannel()
//		select {
//			case <-wch: // waiter done
//				...
//			case <-input:
//			...
//		}
func (w *Waiter) WaitChannel() (ch chan int) {
	ch = make(chan int)
	go func() {
		w.Wait()  // wait for WaitGroup
		close(ch) // close channel (NOTE: doesn't happen if Wait() panics)
	}()
	return
}
