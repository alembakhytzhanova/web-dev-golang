package main

import (
	"io"
	"log"
	"net"
)

func main() {
	// Listen on TCP port 2000 on all interfaces.
	l, err := net.Listen("tcp", "localhost:2000")
	if err != nil {
		log.Fatal(err)

	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.

		go func(c net.Conn) {
			io.Copy(c, c)
			c.Close()
		}(conn)
	}
}
