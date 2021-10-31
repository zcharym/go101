package main

import (
    "log"
    "net"
)

// establishing a TCP connection with Go's standard library
func main() {
    listener, err := net.Listen("tcp", "127.0.0.1:6379")
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

        go func(c net.Conn) {
            defer c.Close()

            conn.Write([]byte("+konw"))

        }(conn)

    }
}
