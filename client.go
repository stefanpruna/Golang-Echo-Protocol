package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// Connect to a host and return the connection object
func connect(host string) net.Conn {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println("Failed to connect to server", err)
		os.Exit(1)
	}
	return conn
}

// Write to the connection
func write(conn net.Conn, message []byte) (int, error) {
	return conn.Write(message)
}

// Read data from socket and insert to channel
func read(conn net.Conn, message chan []byte) {
	reader := bufio.NewReader(conn)
	for {
		bytes, _ := reader.ReadBytes(byte('\n'))
		message <- bytes
	}
}
