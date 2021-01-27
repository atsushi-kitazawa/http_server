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

func parseStartLine(buf bytes.Buffer, req *Request) {
	//fmt.Println("debug[parseStartLine]>")
	startLine := strings.Split(buf.String(), " ")
	req.Method = startLine[0]
	req.Resource = startLine[1]
	req.Version = startLine[2]
}

func parseHeaders(buf bytes.Buffer, req *Request) {
	//fmt.Println("debug[parseHeader]>")
	headers := strings.Split(buf.String(), "\n")
	tmp := make(map[string]string)
	for i := 1; i < len(headers)-1; i++ {
		h := strings.Split(headers[i], ":")
		tmp[h[0]] = h[1]
	}
	req.Headers = tmp
}

func parseBody(buf []byte, req *Request) {
	//fmt.Println("debug[parseBody]>")
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
	parseHeaders(headersBuf, &request)

	if request.Method == "GET" {
	    return request
	}

	// body read.
	bodyBuf := make([]byte, 1024) // must modify read full.
	_, err1 := reader.Read(bodyBuf)
	if err1 != nil {
		panic(err1)
	}
	parseBody(bodyBuf, &request)

	return request
}
