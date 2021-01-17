package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Server running at localhost:9999")

	accept(listener)
}

func accept(listener *net.TCPListener) {
	defer listener.Close()
	for {
		fmt.Println("accepting....")
		conn, err := listener.AcceptTCP()
		if err != nil {
			panic(err)
		}

		go requestHandler(conn)
	}
}

func requestHandler(conn *net.TCPConn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buffer bytes.Buffer
	for {
		ba, _, err := reader.ReadLine()
		if err == io.EOF {
			fmt.Println("EOF")
			break
		} else if err != nil {
			panic(err)
		}
		//
		if "" == string(ba) {
			break
		}
		buffer.Write(ba)
		buffer.Write([]byte("\n"))
	}
	data := buffer.String()
	fmt.Printf("Data> %s", data)

	response(conn)
}

func response(conn *net.TCPConn) {
	var body bytes.Buffer
	body.WriteString("HTTP/1.1 200 OK\n")
	body.WriteString("Content-Type: text/html\n")
	body.WriteString("\n")
	body.WriteString(readRequestFile())
	_, err := conn.Write(body.Bytes())
	if err != nil {
		panic(err)
	}
}

func readRequestFile() string {
	data, err := ioutil.ReadFile("../../index.html")
	if err != nil {
		panic(err)
	}
	return string(data)
}
func printRequest(conn *net.TCPConn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF")
				return
			} else {
				panic(err)
			}
		}

		fmt.Printf("Client> %s", buf)

		n, err = conn.Write(buf[:n])
		if err != nil {
			panic(err)
		}
	}
}
