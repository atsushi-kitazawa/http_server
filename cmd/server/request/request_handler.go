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
	Body     []string
}

func parseRequest(data bytes.Buffer) Request {
	fmt.Println("debug>>", data.String())
	datas := strings.Split(data.String(), "\n")

	var request Request
	// parse start line
	startLine := strings.Split(datas[0], " ")
	request.Method = startLine[0]
	request.Resource = startLine[1]
	request.Version = startLine[2]

	// parse header
	tmp := make(map[string]string)
	var bodyIndex int
	for i := 1; i < len(datas)-1; i++ {
		fmt.Println("debug>", datas[i])
		if len(datas[i]) == 0 {
			// body content from next index
			bodyIndex = i + 1
			break
		}
		h := strings.Split(datas[i], ":")
		tmp[h[0]] = h[1]
	}
	request.Headers = tmp

	// parse body
	// GETのときはこないように
	for i := bodyIndex; i < len(datas)-1; i++ {
		request.Body = append(request.Body, datas[i])
	}

	fmt.Println("method=" + request.Method)
	fmt.Println("resource=" + request.Resource)
	fmt.Println("version=" + request.Version)
	fmt.Println("headers=", request.Headers)
	fmt.Println("body=", request.Body)

	return request
}

func RequestHandler(conn *net.TCPConn) Request {
	reader := bufio.NewReader(conn)

	// start line read.
	var startLineBuf bytes.Buffer
	r, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	startLineBuf.Write(r)
	startLineBuf.Write([]byte("\n"))
	// parse start line

	// headers read.
	var headersBuf bytes.Buffer
	for {
		//fmt.Println("debug1>>>")
		ba, _, err := reader.ReadLine()
		//fmt.Println("debug2>>>", string(ba))
		if err == io.EOF {
			fmt.Println("EOF")
			break
		} else if err != nil {
			panic(err)
		}

		// headers end when appearance of blank.
		if "" == string(ba) {
			break
		}
		headersBuf.Write(ba)
		headersBuf.Write([]byte("\n"))
	}
	// parse headers

	// body read.
	bodyBuf := make([]byte, 1024) // must modify read full.
	//var bodyBuf []byte
	n, err := reader.Read(bodyBuf)
	if err != nil {
		panic(err)
	}
	fmt.Println("debug4>>", n)

	fmt.Println("debug1>>", startLineBuf.String())
	fmt.Println("debug2>>", headersBuf.String())
	fmt.Println("debug3>>", string(bodyBuf))

	// for a moment code
	var request Request
	request.Method = "GET"
	return request

	//data := buffer.String()
	//fmt.Printf("Data> %s", buffer.String())
	//req := parseRequest(buffer)
	//return req
}
