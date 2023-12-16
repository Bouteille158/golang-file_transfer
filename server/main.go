package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("Launched server")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected")
	// Loop, print 1024 times numbers from 1 to 1024 with a wait of 1 second between each print and then send back a message to the client
	for i := 1; i <= 1024; i++ {
		fmt.Println(i)
		_, err := conn.Write([]byte(fmt.Sprintf("%d", i)))
		if err != nil {
			log.Fatal(err)
			break
		}
		time.Sleep(time.Second)
	}
	conn.Write([]byte("I'm done"))

}
