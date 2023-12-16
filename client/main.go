package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Launched client")
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to server")
	for i := 1; i <= 1024; i++ {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(buf[:n])
	}

}
