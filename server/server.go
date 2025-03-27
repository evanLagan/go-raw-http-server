// Core TCP Loop
package server

import (
	"bufio"
	"fmt"
	"net"
)

func Start(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on: ", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept failed: ", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Read failed: ", err)
		return
	}

	fmt.Println("Request Line: ", line)

	// Simple raw response
	body := "Hello from raw TCP server in Go."
	response := "HTTP/1.1 200 OK\r\n" +
		fmt.Sprintf("Content-Length: %d\r\n", len(body)) +
		"Content-Type: text/plain\r\n" +
		"\r\n" +
		body

	conn.Write([]byte(response))

}
