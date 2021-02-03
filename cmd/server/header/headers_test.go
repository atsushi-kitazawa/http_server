package header

import (
	"fmt"
	"os"
	"testing"

	"github.com/atsushi-kitazawa/http_server/cmd/server/request"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
    fmt.Println("before")
    loadContentType()

    code := m.Run()

    fmt.Println("after")
    os.Exit(code)
}

func TestDetermineContentType(t *testing.T) {
    fmt.Println("TestDetermineContentType")

    var req request.Request
    req.Method = "GET"
    req.Version = "HTTP/1.1"
    req.Resource = "/pages/hello.html"
    result := DetermineContentType(req)
    assert.Equal(t, "text/html", result)

    req.Resource = "/hello.css"
    result = DetermineContentType(req)
    assert.Equal(t, "text/css", result)

    req.Resource = "hello.js"
    result = DetermineContentType(req)
    assert.Equal(t, "text/javascript", result)

    req.Resource = "/"
    result = DetermineContentType(req)
    assert.Equal(t, "text/html", result)

    req.Resource = "hello.hoge"
    result = DetermineContentType(req)
    assert.Equal(t, "", result)
}
