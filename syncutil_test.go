package syncutil

import (
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	var waiter Waiter
	waiter.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer waiter.Done()
			time.Sleep(time.Millisecond)
		}()
	}
	timer := time.NewTimer(time.Second)
	channel := waiter.Channel()
	select {
	case <-channel:
		return
	case <-timer.C:
		t.Error("Timed out without Waiter.Channel() being signaled")
	}
}
