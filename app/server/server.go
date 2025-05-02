package server

import (
	"fmt"
	"net"
	"os"
)

func Connect() {
	listen, err := net.Listen("tcp", "0.0.0.0:4221")

	if err != nil {
		fmt.Println("Error accepting connection", err.Error())
		os.Exit(1)
	}

	fmt.Println("Server started on port 4221")

	for {
		connection, err := listen.Accept()

		if err != nil {
			fmt.Println("Error accepting connection", err.Error())
		}

		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	connection.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
}
