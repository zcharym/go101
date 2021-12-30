package behavioral

import (
	"fmt"
	"time"
)

func ExampleNewServer() {
	server := NewServer("127.0.0.1", "80", Protocol("tcp"), Timeout(time.Second*30))
	fmt.Printf("%+v", server)
}

func NewServer(host, port string, options ...Option) *Server {
	server := &Server{
		Host: host,
		Port: port,
	}

	for _, option := range options {
		option(server)
	}

	return server
}
