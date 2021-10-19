package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	const message string = "test\n"

	os.Args = []string{"main", "1234"}

	// start the server
	go main()

	// wait for the server to start
	time.Sleep(1 * time.Second)

	// connect the client to the server
	conn := connect("localhost:1234")

	// create an array of bytes to store the message from the server
	ch := make(chan []byte)

	// start the reader
	go read(conn, ch)

	// write a message to the server
	_, err := write(conn, []byte(message))

	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	} else {
		receivedMessage := <-ch
		receivedString := string(receivedMessage[:])
		if message != receivedString {
			t.Fail()
		}
		fmt.Printf("Client received: %s", receivedMessage)
	}
}
