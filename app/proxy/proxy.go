package proxy

import (
	"io"
	"net"
)

func proxy(client, server net.Conn, bufSize int) error {
	errCh := make(chan error, 2)

	go relay(errCh, client, server, bufSize)
	go relay(errCh, server, client, bufSize)

	return <-errCh
}

func relay(errCh chan error, src io.Reader, dst io.Writer, bufSize int) {
	buf := make([]byte, bufSize)
	_, err := io.CopyBuffer(dst, src, buf)
	errCh <- err
	return
}
