package proxy

import (
	"context"
	"errors"
	"net"

	"github.com/picop-rd/picop-go/propagation"
	picopnet "github.com/picop-rd/picop-go/protocol/net"
	"github.com/picop-rd/proxy/app/entity"
	"github.com/picop-rd/proxy/app/repository"
	"github.com/rs/zerolog/log"
)

type Server struct {
	Env         repository.Env
	Propagate   bool
	DefaultAddr string
	BufSize     int
	closed      bool
	listener    picopnet.Listener
}

func (s *Server) Start(address string) {
	s.closed = false
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal().Err(err).Str("listen address", address).Msg("failed to listen")
	}
	s.listener = picopnet.NewListener(ln)
	defer s.listener.Close()

	log.Info().Msg("starting server")
	for {
		bconn, err := s.listener.AcceptWithPiCoPConn()
		if err != nil {
			if s.closed {
				break
			}
			log.Error().Err(err).Str("listen address", address).Msg("failed to accept")
			continue
		}
		go s.handle(bconn)
	}
}

func (s *Server) Close() {
	log.Info().Msg("proxy shutdown")
	s.closed = true
	s.listener.Close()
}

func (s *Server) handle(clientConn *picopnet.Conn) {
	defer clientConn.Close()

	header, err := clientConn.ReadHeader()
	if err != nil {
		log.Error().
			Err(err).
			Stringer("client local address", clientConn.LocalAddr()).
			Stringer("client remote address", clientConn.RemoteAddr()).
			Msg("invalid PiCoP header")
		return
	}

	envID := header.Get(propagation.EnvIDHeader)

	ctx := context.Background()

	var serverAddr string
	env, err := s.Env.Get(ctx, envID)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			serverAddr = s.DefaultAddr
		} else {
			log.Error().
				Err(err).
				Stringer("client local address", clientConn.LocalAddr()).
				Stringer("client remote address", clientConn.RemoteAddr()).
				Str("env-id", envID).
				Msg("env not found")
			return
		}
	} else {
		serverAddr = env.Destination
	}

	serverConn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Error().
			Err(err).
			Stringer("client local address", clientConn.LocalAddr()).
			Stringer("client remote address", clientConn.RemoteAddr()).
			Str("env-id", envID).
			Str("server address", serverAddr).
			Msg("failed to dial server")
		return
	}
	defer serverConn.Close()

	if s.Propagate {
		_, err := header.WriteTo(serverConn)
		if err != nil {
			log.Error().
				Err(err).
				Stringer("client local address", clientConn.LocalAddr()).
				Stringer("client remote address", clientConn.RemoteAddr()).
				Stringer("server local address", serverConn.LocalAddr()).
				Stringer("server remote address", serverConn.RemoteAddr()).
				Str("env-id", envID).
				Msg("failed to write header")
		}
	}

	err = proxy(clientConn, serverConn, s.BufSize)
	if err != nil {
		log.Error().
			Err(err).
			Stringer("client local address", clientConn.LocalAddr()).
			Stringer("client remote address", clientConn.RemoteAddr()).
			Stringer("server local address", serverConn.LocalAddr()).
			Stringer("server remote address", serverConn.RemoteAddr()).
			Str("env-id", envID).
			Msg("failed to proxy")
	}
}
