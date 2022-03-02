package app

import (
	"bytes"
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type App struct {
	mu      sync.Mutex
	opts    options
	ctx     context.Context
	cancel  func()
	version string
}

func NewApp(opts ...Option) (*App, error) {
	a := new(App)
	a.opts = options{
		ctx:         context.Background(),
		stopTimeout: 10 * time.Second,
		signs:       []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT},
	}
	for _, opt := range opts {
		opt(&a.opts, a)
	}

	a.ctx, a.cancel = context.WithCancel(a.opts.ctx)
	return a, nil
}

func (a *App) Run() error {
	eg, ctx := errgroup.WithContext(a.ctx)
	addr := new(bytes.Buffer)
	for _, server := range a.opts.servers {
		srv := server
		srv.Addr(addr)
		eg.Go(func() error {
			<-ctx.Done()
			sctx, cancel := context.WithTimeout(context.Background(), a.opts.stopTimeout)
			defer cancel()
			return srv.Stop(sctx)
		})
		eg.Go(func() error {
			return srv.Start(ctx)
		})
	}
	log.Printf("Serving %s start with pid: %d and Version: %s", addr.String(), os.Getppid(), a.version)

	c := make(chan os.Signal, 1)
	signal.Notify(c, a.opts.signs...)
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				err := a.Stop()
				if err != nil {
					log.Printf("failed to stop app: %w", err)
					return err
				}
				log.Printf("Serving %s has Done. ", addr.String())
			}
		}
	})
	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil

}

func (a *App) Stop() error {
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}
