package process

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// notify is used to override signal.Notify in unit tests.
var notify = signal.Notify

// HandleSIGTERM cancels the context when the process receives a SIGTERM signal.
// Used to gracefully shut down the application on ctrl+c.
func HandleSIGTERM(ctx context.Context) context.Context {
	ctx, done := context.WithCancel(ctx)

	signals := make(chan os.Signal, 1)
	notify(signals, syscall.SIGTERM)
	go func() {
		<-signals
		done()
	}()

	return ctx
}

// InParallel invokes each of the given functions in a new go routine and
// blocks until they return.
func InParallel(funcs ...func()) {
	var wg sync.WaitGroup

	for _, fn := range funcs {
		wg.Add(1)
		go func(fn func()) {
			defer wg.Done()
			fn()
		}(fn)
	}

	wg.Wait()
}
