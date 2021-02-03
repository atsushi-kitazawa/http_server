package response

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"

	"github.com/atsushi-kitazawa/http_server/cmd/server/header"
	"github.com/atsushi-kitazawa/http_server/cmd/server/request"
)

var indexPage = "/pages/index.html"
var rootDir = getRootDir()

func getRootDir() string {
	d, _ := os.Getwd()
	return d
}

func Response(conn *net.TCPConn, req request.Request) {
	if strings.Contains(req.Resource, "/download") {
	    var body bytes.Buffer
	    body.WriteString("HTTP/1.1 200 OK\n")
	    body.WriteString("Content-Disposition: attachment;filename=\"" + "file" + "\"\n")
	    body.WriteString("\n")
	    body.WriteString(readResource(req))
	    _, err := conn.Write(body.Bytes())
	    if err != nil {
		panic(err)
	    }
	} else {
	    var body bytes.Buffer
	    body.WriteString("HTTP/1.1 200 OK\n")
	    body.WriteString("Content-Type: " + header.DetermineContentType(req) + "\n")
	    body.WriteString("\n")
	    body.WriteString(readResource(req))
	    _, err := conn.Write(body.Bytes())
	    if err != nil {
		panic(err)
	    }
	}
}

func ResponseAuthError(conn *net.TCPConn) {
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
	body.WriteString("HTTP/1.1 404 Not Found\n")
	body.WriteString("Content-Type: text/html\n")
	body.WriteString("\n")
	body.WriteString("<html><body>page not found.</body></html>")
	return body.String()
}
