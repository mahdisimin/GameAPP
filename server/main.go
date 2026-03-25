package main

import (
	"fmt"
	"net"
)

const (
	server   = "127.0.0.1:1234"
	protocol = "tcp"
)

func main() {
	var connection net.Conn
	listener, err := net.Listen(protocol, server)
	if err != nil {
		panic(err)
	}
	fmt.Printf("listening on %s ...\n", listener.Addr())

	if conn, err := listener.Accept(); err != nil {
		panic(err)
	} else {
		connection = conn
	}
	fmt.Printf("recived data from %s ...\n", connection.RemoteAddr())
}
