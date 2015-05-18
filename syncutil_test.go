package syncutil

import (
	"testing"
	"time"
)

func TestWaitChannel(t *testing.T) {
	var waiter Waiter
	waiter.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer waiter.Done()
			time.Sleep(time.Millisecond)
		}()
	}
	timer := time.NewTimer(time.Second)
	waitch := waiter.WaitChannel()
	select {
	case <-waitch:
		return
	case <-timer.C:
		t.Error("Timed out without Waiter.Channel() being signaled")
	}
}
