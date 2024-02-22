package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"net"
	"os"
	"path/filepath"
)

var payloadFolder = "./payload/"
var packetSize = 1024
var serverAddress = "localhost:8080"

func main() {
	files, err := os.ReadDir(payloadFolder)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		fmt.Printf("Sending %s\n", file)
		sendFile(file.Name())
	}
}

func sendFile(filename string) {
	fmt.Println("Launched client")

	file, err := os.Open(payloadFolder + filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Obtenez la taille du fichier
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	size := fileInfo.Size()

	// Get file extension
	fileExtension := filepath.Ext(fileInfo.Name())

	fmt.Printf("Size: %d\n", size)
	fmt.Printf("Extension: %s\n", fileExtension)

	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to server")

	packetCount := int64(math.Ceil(float64(size) / float64(packetSize)))
	fmt.Printf("Packet count: %d\n", packetCount)

	packetCountBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(packetCountBytes, uint64(packetCount))

	extensionSizeBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(extensionSizeBytes, uint64(len(fileExtension)))

	fileExtensionBytes := []byte(fileExtension)

	fmt.Printf("Packet count bytes: %v\n", packetCountBytes)
	fmt.Printf("Extension size bytes: %v\n", extensionSizeBytes)
	fmt.Printf("Extension bytes: %v\n", fileExtensionBytes)

	fmt.Println("Sending payload size")
	_, err = conn.Write(packetCountBytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Sending file extension size")
	_, err = conn.Write(extensionSizeBytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Sending file extension")
	_, err = conn.Write(fileExtensionBytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create buffer to read file
	buf := make([]byte, packetSize)

	for i := int64(1); i <= packetCount; i++ {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}

		// fmt.Printf("Read %d bytes\n", n)
		// fmt.Printf("Read %d \n", buf[:n])

		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println(err)
			break
		}
		//Print progress bar
		fmt.Printf("%d/%d\n", i, packetCount)
		// Print percentage
		fmt.Printf("%d%%\n", int(float64(i)/float64(packetCount)*100))

		// Go back 2 lines and erase them
		if i < packetCount {
			fmt.Printf("\033[2A")
		}
	}

	fmt.Println("Done sending payload")

	message := make([]byte, 1024)
	n, err := conn.Read(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(message[:n]))

}
