package transport

import (
	"bytes"
	"context"
)

//ServerInterface is transport server
type ServerInterface interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Addr(b *bytes.Buffer)
}
