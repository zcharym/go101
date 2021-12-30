package behavioral

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	serverBuilder := &ServerBuilder{}

	server := serverBuilder.
		create("127.0.0.1", "80").
		withTimeout(time.Second * 30).
		withProtocol("tcp").
		build()

	expectedServer := Server{
		Host: "127.0.0.1",
		Port: "80",
		Conf: &Config{
			Protocol: "tcp",
			Timeout:  time.Second * 30,
		},
	}

	assert.EqualValues(t, expectedServer, server)
}
