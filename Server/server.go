package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

var port = "0.0.0.0:9001"
var serverMsg = []byte("Server: ")

func echo(conn net.Conn) {
	r := bufio.NewReader(conn)
	var finished = false
	for !finished {
		line, err := r.ReadBytes(byte('\n'))
		switch err {
		case nil:
			fmt.Printf("New text requested from %s\n", conn.LocalAddr())
			break
		case io.EOF:
		default:
			fmt.Println("Connection closed: ", err)
			finished = true
		}
		conn.Write(line)
	}
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}

	fmt.Println("Server started at " + port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("ERROR", err)
			continue
		}
		fmt.Printf("Listen a client in %s\n", conn.LocalAddr())
		go echo(conn)
	}
}
