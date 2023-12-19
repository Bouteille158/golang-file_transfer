package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

var outputFolder = "./reception"
var packetSize = 1024

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

	payloadSizeBytes := make([]byte, 8)
	_, err := conn.Read(payloadSizeBytes)
	if err != nil {
		log.Fatal("Error while reading payload size:\n", err)
		return
	}

	extensionSizeBytes := make([]byte, 8)
	_, err = conn.Read(extensionSizeBytes)
	if err != nil {
		log.Fatal("Error while reading extension size:\n", err)
		return
	}

	extensionSize := int64(binary.BigEndian.Uint64(extensionSizeBytes))

	fileExtensionBytes := make([]byte, extensionSize)
	_, err = conn.Read(fileExtensionBytes)
	if err != nil {
		log.Fatal("Error while reading file extension:\n", err)
		return
	}

	fmt.Printf("Payload size bytes: %v\n", payloadSizeBytes)
	// Convert payload size bytes to int64
	payloadSize := int64(binary.BigEndian.Uint64(payloadSizeBytes))
	fmt.Printf("Payload size: %d\n", payloadSize)

	outputFileName := time.Now().Format("2006-01-02_15-04-05") + "_" + randomString(10) + "." + string(fileExtensionBytes)

	// Create output folder if it doesn't exist
	if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
		err = os.Mkdir(outputFolder, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	//Create output file
	_, err = os.Create(outputFolder + "/" + outputFileName)
	if err != nil {
		log.Fatal(err)
	}

	for i := int64(1); i <= payloadSize; i++ {
		buf := make([]byte, packetSize)
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
			break
		}
		fmt.Print(buf[:n])

		f, err := os.OpenFile(outputFolder+"/"+outputFileName, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err := f.Write(buf[:n]); err != nil {
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("\nDone reading payload")
	conn.Write([]byte("I'm done"))

}

func randomString(length int) string {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = byte(rand.Intn(26) + 97) // Génère un nombre aléatoire entre 97 ('a') et 122 ('z')
	}
	return string(bytes)
}
