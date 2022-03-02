package app

import (
	"bookstore/pkg/transport"
	"context"
	"os"
	"time"
)

type Option func(o *options, app *App)

type options struct {
	id          string
	name        string
	ctx         context.Context
	stopTimeout time.Duration
	signs       []os.Signal
	servers     []transport.ServerInterface
}

func ID(id string) Option {
	return func(o *options, app *App) {
		o.id = id
	}
}

func Name(name string) Option {
	return func(o *options, app *App) {
		o.name = name
	}
}

func Context(ctx context.Context) Option {
	return func(o *options, app *App) {
		o.ctx = ctx
	}
}

func Signal(signs ...os.Signal) Option {
	return func(o *options, app *App) {
		o.signs = signs
	}
}

func Server(servers ...transport.ServerInterface) Option {
	return func(o *options, app *App) {
		o.servers = servers
	}
}

func StopTimeout(t time.Duration) Option {
	return func(o *options, app *App) {
		o.stopTimeout = t
	}
}

func Version(version string) Option {
	return func(o *options, app *App) {
		app.version = version
	}
}
