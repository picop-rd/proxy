package proxy

import (
	"io"
	"net"
)

func proxy(client, server net.Conn) error {
	errCh := make(chan error, 2)

	go relay(errCh, client, server)
	go relay(errCh, server, client)

	return <-errCh
}

func relay(errCh chan error, src io.Reader, dst io.Writer) {
	_, err := io.Copy(dst, src)
	errCh <- err
	return
}
