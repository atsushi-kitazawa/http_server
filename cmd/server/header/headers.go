package header

import (
	"strings"

	"github.com/atsushi-kitazawa/http_server/cmd/server/request"
)

// [key, value] = [extension, content-type]
var content_type map[string]string = loadContentType()

func loadContentType() map[string]string {
    m := make(map[string]string)
    m[".html"] = "text/html"
    m[".js"] = "text/javascript"
    m[".css"] = "text/css"
    return m
}

func DetermineContentType(req request.Request) string {
    if req.Resource == "/" {
	return "text/html"
    }

    periodPos := strings.LastIndex(req.Resource, ".")
    if periodPos == -1 {
	return ""
    }
    extension := req.Resource[periodPos:]
    return content_type[extension]
}

func IsMultipartFormData(headers map[string]string) bool {
    contentType := headers["Content-Type"]
    if strings.Contains(contentType, "multipart/form-data") {
	return true
    }
    return false
}
