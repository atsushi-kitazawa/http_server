package response

import (
	"bytes"
	"github.com/atsushi-kitazawa/http_server/cmd/server/request"
	"io/ioutil"
	"net"
)

func Response(conn *net.TCPConn, req request.Request) {
	var body bytes.Buffer
	body.WriteString("HTTP/1.1 200 OK\n")
	body.WriteString("Content-Type: text/html\n")
	body.WriteString("\n")
	body.WriteString(readResource(req))
	_, err := conn.Write(body.Bytes())
	if err != nil {
		panic(err)
	}
}

func readResource(req request.Request) string {
	if req.Resource != "/" {
		data, err := ioutil.ReadFile("../../" + req.Resource)
		if err != nil {
			panic(err)
		}
		return string(data)
	} else {
		data, err := ioutil.ReadFile("../../index.html")
		if err != nil {
			panic(err)
		}
		return string(data)
	}
}
