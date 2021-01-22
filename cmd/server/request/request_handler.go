package request

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"strings"
)

type Request struct {
	Method   string
	Resource string
	Version  string
	Headers  map[string]string
	Body     string
}

func parseRequest(data bytes.Buffer) Request {
	datas := strings.Split(data.String(), "\n")

	var request Request
	// parse start line
	startLine := strings.Split(datas[0], " ")
	request.Method = startLine[0]
	request.Resource = startLine[1]
	request.Version = startLine[2]

	// parse header
	tmp := make(map[string]string)
	for i := 1; i < len(datas)-1; i++ {
		h := strings.Split(datas[i], ":")
		tmp[h[0]] = h[1]
	}
	request.Headers = tmp

	// parse body

	//fmt.Println(request.Method)
	//fmt.Println(request.Resource)
	//fmt.Println(request.Version)
	//fmt.Println(request.Headers)

	return request
}

func RequestHandler(conn *net.TCPConn) Request {
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
	req := parseRequest(buffer)

	return req
}
