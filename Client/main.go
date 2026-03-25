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
	if conn, err := net.Dial(protocol, server); err != nil {
		panic(err)
	} else {
		connection = conn
	}
	defer connection.Close()
	fmt.Printf("coonecting to server %v\n", server)
}
