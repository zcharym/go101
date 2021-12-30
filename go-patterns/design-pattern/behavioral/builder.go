package behavioral

import (
	"time"
)

type Server struct {
	Host string
	Port string
	Conf *Config
}

// Config is optional fields for Server
type Config struct {
	Protocol string
	Timeout  time.Duration
}

type ServerBuilder struct {
	Server
}

func (s *ServerBuilder) create(host, port string) *ServerBuilder {
	s.Server.Port = port
	s.Server.Host = host
	return s
}

func (s *ServerBuilder) withProtocol(proto string) *ServerBuilder {
	if s.Server.Conf == nil {
		s.Server.Conf = &Config{}
	}
	s.Server.Conf.Protocol = proto
	return s
}

func (s *ServerBuilder) withTimeout(timeout time.Duration) *ServerBuilder {
	if s.Server.Conf == nil {
		s.Server.Conf = &Config{}
	}
	s.Server.Conf.Timeout = timeout
	return s
}

func (s *ServerBuilder) build() Server {
	return s.Server
}
