package http_server

import (
	"context"
	"net"
	"net/http"
	"time"
)

const (
	DefaultReadTimeout     = 5 * time.Second
	DefaultWriteTimeout    = 5 * time.Second
	DefaultShutdownTimeout = 5 * time.Second
)

type Server struct {
	httpServer      *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func New(handler http.Handler, port string) *Server {
	addr := net.JoinHostPort("", port)

	httpServer := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  DefaultReadTimeout,
		WriteTimeout: DefaultWriteTimeout,
	}

	s := &Server{
		httpServer:      httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: DefaultShutdownTimeout,
	}
	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.httpServer.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}
