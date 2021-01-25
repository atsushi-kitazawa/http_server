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
	//fmt.Println("debug>>", data.String())
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
		//request.Body = append(request.Body, datas[i])
	}

	fmt.Println("method=" + request.Method)
	fmt.Println("resource=" + request.Resource)
	fmt.Println("version=" + request.Version)
	fmt.Println("headers=", request.Headers)
	fmt.Println("body=", request.Body)

	return request
}

func parseStartLine(buf bytes.Buffer, req *Request) {
	fmt.Println("debug[parseStartLine]>")
	startLine := strings.Split(buf.String(), " ")
	req.Method = startLine[0]
	req.Resource = startLine[1]
	req.Version = startLine[2]
}

func parseHeaders(buf bytes.Buffer, req *Request) {
	fmt.Println("debug[parseHeader]>")
	headers := strings.Split(buf.String(), "\n")
	tmp := make(map[string]string)
	for i := 1; i < len(headers)-1; i++ {
		h := strings.Split(headers[i], ":")
		tmp[h[0]] = h[1]
	}
	req.Headers = tmp
}

func parseBody(buf []byte, req *Request) {
	fmt.Println("debug[parseBody]>")
	req.Body = string(buf)
}

func RequestHandler(conn *net.TCPConn) Request {
	reader := bufio.NewReader(conn)

	// result request of parsing
	var request Request

	// start line read.
	var startLineBuf bytes.Buffer
	r, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	startLineBuf.Write(r)
	startLineBuf.Write([]byte("\n"))
	// parse start line
	parseStartLine(startLineBuf, &request)

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
	parseHeaders(headersBuf, &request)

	// body read.
	bodyBuf := make([]byte, 1024) // must modify read full.
	//var bodyBuf []byte
	_, err1 := reader.Read(bodyBuf)
	if err1 != nil {
		panic(err1)
	}
	// parse body
	parseBody(bodyBuf, &request)

	//data := buffer.String()
	//fmt.Printf("Data> %s", buffer.String())
	//req := parseRequest(buffer)
	return request
}
