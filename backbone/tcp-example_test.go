package main

import (
    "net"
    "testing"
)

func TestListener(t *testing.T)  {
     listener, err := net.Listen("tcp","127.0.0.1:0") // empty or zero, go will randomly assign a port number
     if err != nil {
         t.Fatal(err)
     }
     defer func() {
         _ = listener.Close()
     }()
     t.Logf("bound to %q", listener.Addr())
}
