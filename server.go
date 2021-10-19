/*
Echo server implementation in Golang
https://datatracker.ietf.org/doc/html/rfc862
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

// Check if a given port is valid
func isPortValid(port int) bool {
	if port > 1 && port < 65535 {
		return true
	}
	return false
}

// Listen on a given port
func listen(port string) {

	// Createa listener
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		fmt.Println("Failed to create listener, err:", err)
		os.Exit(1)
	}

	// Accept connection in an infinite loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection, err:", err)
		}

		// Handle the connection as per protocol
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("got connection")
	defer conn.Close()

	//create a reader for the connection
	reader := bufio.NewReader(conn)
	for {
		// read client request data
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("Failed to read data, err:", err)
			}
		}

		fmt.Printf("Server received message: %s", bytes)

		// Write the message back to the client
		conn.Write(bytes)
	}
}

func main() {

	// Check the number of arguments
	if len(os.Args) < 1 {
		fmt.Println("Server needs to be ran with a port to listen to as an argument")
		os.Exit(1)
	}

	// Get the port from the command line arguments
	port := string(os.Args[1])
	port_integer, err := strconv.Atoi(port)

	//validate the port
	if err != nil {
		fmt.Println("Port needs to be a valid number:", err)
		os.Exit(1)
	}

	if !isPortValid(port_integer) {
		fmt.Println("Given port is not a valid port")
		os.Exit(1)
	}

	// listen for connection on the port
	listen(port)
}
