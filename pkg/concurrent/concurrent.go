package concurrent

import (
	"context"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

// SetupFunc is described in Setup.
type SetupFunc func(Ctx) (interface{}, error)

// Serve runs given services in parallel until either ctx.Done or any
// service exits, then it call cancel and wait until all services will
// exit.
// Returns error of first service which returned non-nil error, if any.
func Serve(ctx Ctx, cancel func(), services ...func(Ctx) error) (err error) {
	errc := make(chan error)
	for _, service := range services {
		service := service
		go func() { errc <- service(ctx) }()
	}
	for range services {
		if err == nil {
			err = <-errc
		} else {
			<-errc
		}
		cancel()
	}
	return err
}
