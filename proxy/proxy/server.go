package proxy

import (
	"context"
	"errors"
	"net"

	bcopnet "github.com/hiroyaonoe/bcop-go/protocol/net"
	"github.com/hiroyaonoe/bcop-proxy/proxy/entity"
	"github.com/hiroyaonoe/bcop-proxy/proxy/repository"
	"github.com/rs/zerolog/log"
)

type Server struct {
	Env                         repository.Env
	GetEnvIDFromHeaderValueFunc func(string) (string, error)
	Propagate                   bool
	DefaultAddr                 string
}

func (s *Server) Start(address string) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal().Err(err).Str("listen address", address).Msg("failed to listen")
	}
	bln := bcopnet.NewListener(ln)
	defer bln.Close()

	log.Info().Msg("starting server")
	for {
		bconn, err := bln.AcceptWithBCoPConn()
		if err != nil {
			log.Fatal().Err(err).Str("listen address", address).Msg("failed to accept")
		}

		go s.handle(bconn)
	}
}

func (s *Server) handle(clientConn *bcopnet.Conn) {
	defer clientConn.Close()

	header, err := clientConn.ReadHeader()
	if err != nil {
		log.Error().
			Err(err).
			Stringer("client local address", clientConn.LocalAddr()).
			Stringer("client remote address", clientConn.RemoteAddr()).
			Msg("invalid BCoP header")
		return
	}

	envID, err := s.GetEnvIDFromHeaderValueFunc(header.Get())
	if err != nil {
		log.Error().
			Err(err).
			Stringer("client local address", clientConn.LocalAddr()).
			Stringer("client remote address", clientConn.RemoteAddr()).
			Msg("failed to parse BCoP header")
		return
	}

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

	err = proxy(clientConn, serverConn)
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
