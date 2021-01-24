package response

import (
	"bytes"
	"fmt"
	"github.com/atsushi-kitazawa/http_server/cmd/server/request"
	"io/ioutil"
	"net"
	"os"
)

var indexPage = "/pages/index.html"
var rootDir = getRootDir()

func getRootDir() string {
	d, _ := os.Getwd()
	return d
}

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
		data, err := ioutil.ReadFile(rootDir + req.Resource)
		if err != nil {
			fmt.Println(err)
			return notFoundResponse()
		}
		return string(data)
	} else {
		data, err := ioutil.ReadFile(rootDir + indexPage)
		if err != nil {
			panic(err)
		}
		return string(data)
	}
}

func notFoundResponse() string {
	var body bytes.Buffer
	body.WriteString("HTTP/1.1 404 Not Found")
	body.WriteString("Content-Type: text/html")
	body.WriteString("\n")
	body.WriteString("<html><body>page not found.</body></html>")
	return body.String()
}
