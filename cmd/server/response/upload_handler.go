package response

import (
    "net"
    "bytes"

    "github.com/atsushi-kitazawa/http_server/cmd/server/request"
)

func ResponseUpload(conn *net.TCPConn, req request.Request) {
    var res bytes.Buffer
    res.WriteString("HTTP/1.1 200 OK\n")
    res.WriteString("Content-Type: text/html\n")
    //res.WriteString("application/octet-stream\n")
    res.WriteString("\n")
    res.WriteString("<html><body> upload ok.</body></html>")
    _, err := conn.Write(res.Bytes())
    if err != nil {
	panic(err)
    }
    return
}
