package main

import (
	"fmt"
	"time"
)

// no functional options used

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

type Option func(rs *RealServer)

type RealServer struct {
	addr string

	timeout   time.Duration // Default no timeout
	tlsConfig *tlsConfig
}

type tlsConfig struct {
}

// Timeout configures a max length of idle connection in RealServer.
func Timeout(timeout time.Duration) Option {
	return func(s *RealServer) {
		s.timeout = timeout
	}
}

func TLSConfig(tlsConfig *tlsConfig) Option {
	return func(s *RealServer) {
		s.tlsConfig = tlsConfig
	}
}

func NewRealServer(addr string, opts ...Option) *RealServer {
	server := &RealServer{addr: addr}

	// apply the list of options to the RealServer
	for _, opt := range opts {
		opt(server)
	}

	return server
}

func main() {
	newserver := NewServer("127.0.0.1:8080")
	fmt.Println(newserver)

	// Functional options in use, easy to understand
	server := NewRealServer("127.0.0.1:8080", Timeout(30*time.Second), TLSConfig(&tlsConfig{}))
	fmt.Println(server)
	Timeout(999 * time.Hour)(server) // Potential race condition, we don't want to be able to change the timeout once already created
}
