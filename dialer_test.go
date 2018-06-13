package http_dialer

import (
	"context"
	"net"
	"testing"
	"time"
)

func TestCloseAll(t *testing.T) {
	closed := make(chan struct{})
	dialFn := func(ctx context.Context, network, address string) (net.Conn, error) {
		return closeOnlyConn{onClose: func() { closed <- struct{}{} }}, nil
	}
	dialer := NewDialer(dialFn)

	const numConns = 10

	// Outer loop to ensure Dialer is re-usable after CloseAll.
	for i := 0; i < 5; i++ {
		for j := 0; j < numConns; j++ {
			if _, err := dialer.Dial("", ""); err != nil {
				t.Fatal(err)
			}
		}
		dialer.CloseAll()
		for j := 0; j < numConns; j++ {
			select {
			case <-closed:
			case <-time.After(time.Second):
				t.Fatalf("iteration %d: 1s after CloseAll only %d/%d connections closed", i, j, numConns)
			}
		}
	}
}

type closeOnlyConn struct {
	net.Conn
	onClose func()
}

func (c closeOnlyConn) Close() error {
	go c.onClose()
	return nil
}
