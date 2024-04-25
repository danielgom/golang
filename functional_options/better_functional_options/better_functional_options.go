package main

import (
	"fmt"
	"time"
)

type ServerOption interface {
	// apply is unexported
	// only current package can implement this interface
	applyServer(s *Server)
}

type ClientOption interface {
	applyClient(c *Client)
}

type Option interface {
	ServerOption
	ClientOption
}

// Timeout option to Server and Client
type Timeout time.Duration

func (t withTimeout) applyServer(s *Server) {
	s.timeout = time.Duration(t.timeout)
}

func (t withTimeout) applyClient(c *Client) {
	c.timeout = time.Duration(t.timeout)
}

type withTimeout struct {
	timeout Timeout
}

func WithTimeout(timeout time.Duration) Option {
	if timeout == 0 {
		return withTimeout{timeout: 0}
	}
	return withTimeout{timeout: Timeout(timeout)}
}

type Logger string

type withLogger struct {
	logger Logger
}

func (i withLogger) applyServer(s *Server) {
	s.logger = i.logger
}

func (i withLogger) applyClient(c *Client) {
	c.logger = i.logger
}

func WithLogger(logger Logger) Option {
	return withLogger{logger: logger}
}

type withPort struct {
	port *int
}

func (p withPort) applyServer(s *Server) {
	if *p.port < 0 {
		panic("port cannot be negative")
	}
	s.port = *p.port
}

func WithPort(port int) ServerOption {
	return withPort{port: &port}
}

type Server struct {
	addr string

	timeout time.Duration
	logger  Logger
	port    int
}

func NewServer(addr string, serverOptions ...ServerOption) *Server {
	// Defaults
	server := &Server{
		addr: addr,
		port: 8080, // Default port
	}

	for _, option := range serverOptions {
		option.applyServer(server)
	}

	return server
}

type Client struct {
	timeout time.Duration
	logger  Logger
}

func NewClient(serverOptions ...ClientOption) *Client {
	// Defaults
	client := &Client{}

	for _, option := range serverOptions {
		option.applyClient(client)
	}

	return client
}

func main() {
	ss := NewServer("127.0.0.1", WithLogger("server"), WithTimeout(time.Hour*10), WithPort(10))
	fmt.Println(ss)
}
