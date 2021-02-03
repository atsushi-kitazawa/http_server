package response

import (
	"bytes"
	"errors"
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
	// first, read request resource for decide status code
	body, err := readResource(req)

	// 404 Not Found Response.
	if err != nil {
	    var res bytes.Buffer
	    res.WriteString(notFoundResponse())
	    _, err := conn.Write(res.Bytes())
	    if err != nil {
		panic(err)
	    }
	    return
	}

	// 200 OK Response
	contentType := header.DetermineContentType(req)
	fmt.Println(contentType)
	if contentType == "" {
	    var res bytes.Buffer
	    res.WriteString("HTTP/1.1 200 OK\n")
	    res.WriteString("Content-Disposition: attachment;filename=\"" + getFileName(req) + "\"\n")
	    //res.WriteString("application/octet-stream\n")
	    res.WriteString("\n")
	    res.WriteString(body)
	    _, err := conn.Write(res.Bytes())
	    if err != nil {
		panic(err)
	    }
	    return
	}

	var res bytes.Buffer
	res.WriteString("HTTP/1.1 200 OK\n")
	res.WriteString("Content-Type: " + header.DetermineContentType(req) + "\n")
	res.WriteString("\n")
	res.WriteString(body)
	_, err = conn.Write(res.Bytes())
	if err != nil {
	    panic(err)
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

func readResource(req request.Request) (string, error) {
	if req.Resource != "/" {
		data, err := ioutil.ReadFile(rootDir + req.Resource)
		if err != nil {
			fmt.Println(err)
			return "", errors.New("Not Found")
		}
		return string(data), nil
	} else {
		data, err := ioutil.ReadFile(rootDir + indexPage)
		if err != nil {
			panic(err)
		}
		return string(data), nil
	}
}

func getFileName(req request.Request) string {
    resource := req.Resource
    index := strings.LastIndex(resource, "/")
    return resource[index+1:]
}

func notFoundResponse() string {
	var body bytes.Buffer
	body.WriteString("HTTP/1.1 404 Not Found\n")
	body.WriteString("Content-Type: text/html\n")
	body.WriteString("\n")
	body.WriteString("<html><body>page not found.</body></html>")
	return body.String()
}
