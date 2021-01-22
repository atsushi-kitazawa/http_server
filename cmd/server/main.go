package main

import (
	"fmt"
	"github.com/atsushi-kitazawa/http_server/cmd/server/request"
	"github.com/atsushi-kitazawa/http_server/cmd/server/response"
	"io"
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

		go handler(conn)
	}
}

func handler(conn *net.TCPConn) {
	defer conn.Close()
	req := request.RequestHandler(conn)
	response.Response(conn, req)
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
