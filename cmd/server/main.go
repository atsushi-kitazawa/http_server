package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"bytes"

	"github.com/atsushi-kitazawa/http_server/cmd/server/request"
	"github.com/atsushi-kitazawa/http_server/cmd/server/response"
	"github.com/atsushi-kitazawa/http_server/configs"
)

var conf configs.Configuration

func init() {
    // load configuration file
    configs.Load(&conf)
}

func main() {
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
	if isAuth(&req) {
	    fmt.Println("debug:isAuth true")
	    checkAuthHeader(conn, &req)
	    return
	}
	response.Response(conn, req)
}

// chech request resource is required authorization.
func isAuth(req *request.Request) bool {
   if strings.Contains(req.Resource, conf.Auth.Location) {
       return true
   }
   return false
}

func checkAuthHeader(conn *net.TCPConn, req *request.Request) {
	// request has Ahthorization header.
	if cred, ok := req.Headers["Authorization"]; ok {
	    if strings.Trim(cred, "Basic ") == "dGVzdDp0ZXN0" {
		fmt.Println("debug:cred match")
		response.Response(conn, *req)
	    } else {
		fmt.Println("debug:cred not match")
		var body bytes.Buffer
		body.WriteString("HTTP/1.1 401 Unauthorized\n")
		body.WriteString("Content-Type: text/html\n")
		body.WriteString("\n")
		body.WriteString("<html><body>Unauthorized</body></html>")
		_, err := conn.Write(body.Bytes())
		if err != nil {
		    panic(err)
		}
	    }
	    return
	}

	// request has no Authorization header.
	var body bytes.Buffer
	body.WriteString("HTTP/1.1 401 Unauthorized\n")
	body.WriteString(fmt.Sprintf("WWW-Authenticate: %s realm=basic authentication\n", conf.Auth.Type))
	body.WriteString("\n")
	_, err := conn.Write(body.Bytes())
	if err != nil {
		panic(err)
	}
}

func chechAuth(req *request.Request) bool {
    if _, ok := req.Headers["Authorization"]; !ok {
	return false
    }
    cred := req.Headers["Authorization"]
    fmt.Println("debug:cred>", cred)
    if cred == "dGVzdDp0ZXN0" {
	return true
    }
    return false
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
