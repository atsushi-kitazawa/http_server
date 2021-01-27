package main

import (
	"fmt"
	"io"
	"net"

	"github.com/atsushi-kitazawa/http_server/cmd/server/request"
	"github.com/atsushi-kitazawa/http_server/cmd/server/response"
	"github.com/atsushi-kitazawa/http_server/configs"
)

func main() {
	// load configuration file
	var conf configs.Configuration
	configs.Load(&conf)
	//fmt.Println("ip>", conf.Ip)
	//fmt.Println("port>", conf.Port)

	tcpAddr, err := net.ResolveTCPAddr("tcp", conf.Ip + ":" + conf.Port)
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server running at %s:%s\n", conf.Ip, conf.Port)

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
