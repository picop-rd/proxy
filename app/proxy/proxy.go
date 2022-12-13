package proxy

import (
	"io"
	"net"

	"golang.org/x/sync/errgroup"
)

func proxy(client, server net.Conn) error {
	var eg errgroup.Group

	eg.Go(func() error { return relay(client, server) })
	eg.Go(func() error { return relay(server, client) })

	return eg.Wait()
}

func relay(src io.Reader, dst io.Writer) error {
	_, err := io.Copy(dst, src)
	return err
}
