package main

import (
	"log"
	"net"
	"time"
)

// establishing a TCP connection with Go's standard library
func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = listener.Close()
	}()
	log.Printf("bound to %q", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer func() {
		_ = c.Close()
	}()
	log.Printf("new connection from %q", c.RemoteAddr())
	for {
		err := c.SetDeadline(time.Now().Add(time.Second * 1))
		if err != nil {
			log.Print(err)
			return
		}

		_, err = c.Write([]byte("world\n"))
		if err != nil {
			log.Print(err)
			return
		}
	}
}
