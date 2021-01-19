package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"strings"
)

type Request struct {
	method   string
	resource string
	version  string
	headers  map[string]string
	body     string
}

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

func parseRequest(data bytes.Buffer) {
	datas := strings.Split(data.String(), "\n")

	var request Request
	// parse start line
	startLine := strings.Split(datas[0], " ")
	request.method = startLine[0]
	request.resource = startLine[1]
	request.version = startLine[2]

	// parse header
	tmp := make(map[string]string)
	for i := 1; i < len(datas)-1; i++ {
		h := strings.Split(datas[i], ":")
		tmp[h[0]] = h[1]
	}
	request.headers = tmp

	// parse body

	//fmt.Println(request.method)
	//fmt.Println(request.resource)
	//fmt.Println(request.version)
	//fmt.Println(request.headers)
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
	//data := buffer.String()
	//fmt.Printf("Data> %s", data)
	parseRequest(buffer)

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
