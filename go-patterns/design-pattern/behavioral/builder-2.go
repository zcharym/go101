package behavioral

import (
	"time"
)

type Option func(*Server)

func Timeout(duration time.Duration) Option {
	return func(server *Server) {
		server.Conf.Timeout = duration
	}
}

func Protocol(proto string) Option {
	return func(server *Server) {
		server.Conf.Protocol = proto
	}
}
