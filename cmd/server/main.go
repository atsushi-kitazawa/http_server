package main

import (
	"fmt"
	"io"
	"net"

	//"strings"
	//"bytes"

	"github.com/atsushi-kitazawa/http_server/cmd/server/auth"
	"github.com/atsushi-kitazawa/http_server/cmd/server/enviroment"
	"github.com/atsushi-kitazawa/http_server/cmd/server/request"
	"github.com/atsushi-kitazawa/http_server/cmd/server/response"
	"github.com/atsushi-kitazawa/http_server/configs"
)

var conf *configs.Configuration

func main() {
	// parse arguments
	args := enviroment.GetArgs()

	// load configuration file
	conf = configs.Load(args.ConfFile)

	// init auth module
	auth.InitAuthModule()

	// main
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

	if !auth.IsAuthRequireResource(&req) {
	    response.Response(conn, req)
	    return
	}

	if !auth.CheckAuth(&req) {
	   response.ResponseAuthError(conn)
	   return
	}

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
