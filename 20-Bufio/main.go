package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// reader := bufio.NewReader(strings.NewReader("Hello, Bufio packageee!\nHow are you doing?\n"))

	// // Reading byte slice => Read
	// data := make([]byte, 20)
	// // Read() will read data upto the limit
	// n, err := reader.Read(data)
	// if err != nil {
	// 	fmt.Println("error: while reading the data into byte")
	// 	return
	// }
	// fmt.Printf("Read %d bytes %s:\n", n, data[:n])

	// // ReadString(delim)
	// line, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("error reading string:", err)
	// 	return
	// }
	// fmt.Println(line)

	// // ReadByte()
	// byte1, _ := reader.ReadByte()
	// byte2, _ := reader.ReadByte()
	// byte3, _ := reader.ReadByte()
	// fmt.Println(string(byte1))
	// fmt.Println(string(byte2))
	// fmt.Println(string(byte3))

	// // ReadBytes(\n)
	// b, _ := reader.ReadBytes('\n')
	// fmt.Println(b)
	// fmt.Println(string(b))

	// Bufio Writer
	writer := bufio.NewWriter(os.Stdout)

	// Writing byte slice
	data := []byte("Hello, bufio package!\n")
	n, err := writer.Write(data)
	if err != nil {
		fmt.Println("error: while writing into dest file")
		return
	}
	fmt.Printf("Write %d bytes ", n)

	// WriteString()
	str := "Welcome to Golang course from Basic to gRPC"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("error: while writing string on os.Stdout")
		return
	}
	fmt.Printf("wrote %d bytes\n", n)

	// Flush the buffer to ensure that all the data is written to os.Stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println("error: while flusing writer")
		return
	}
}
