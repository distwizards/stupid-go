package process

import (
	"context"
	"os"
	"os/signal"
	"testing"
)

func TestHandleSIGTERM(t *testing.T) {
	signals, done := fakeNotifySignal()
	defer done()

	ctx := HandleSIGTERM(context.Background())

	if ctx.Err() != nil {
		t.Fatalf("expected context to not be canceled, got error: %v", ctx.Err())
	}

	signals <- nil // Send a signal
	<-ctx.Done()   // Context is canceled

	if ctx.Err() != context.Canceled {
		t.Fatalf("expected context to be canceled, got error: %v", ctx.Err())
	}
}

func TestInParallel(t *testing.T) {
	ch := make(chan bool)
	defer close(ch)

	fn := func() {
		ch <- true // Notify that fn was called
		<-ch       // Block
	}

	go func() {
		InParallel(fn) // Invoke fn
		ch <- true     // Notify that InParallel returned
	}()

	<-ch       // Wait for fn to be called
	ch <- true // Unblock fn
	<-ch       // Wait for InParallel to return
}

func fakeNotifySignal() (chan os.Signal, func()) {
	signals := make(chan os.Signal)

	// Forward events from our signal channel to the
	// caller's signal channel
	notify = func(c chan<- os.Signal, sig ...os.Signal) {
		go func() {
			for signal := range signals {
				c <- signal
			}
		}()
	}

	cleanup := func() {
		notify = signal.Notify

		// Don't leak the signal forward loop goroutine
		close(signals)
	}

	return signals, cleanup
}
