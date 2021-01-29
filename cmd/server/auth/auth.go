package auth

import (
	"strings"

	"github.com/atsushi-kitazawa/http_server/cmd/server/request"
	"github.com/atsushi-kitazawa/http_server/configs"
)

// chech request resource is required authorization.
func IsAuthRequireResource(req *request.Request, conf *configs.Configuration) bool {
   if strings.Contains(req.Resource, conf.Auth.Location) {
       return true
   }
   return false
}

// chech Authorization header value is valid (support only Basic Auth)
func CheckAuth(req *request.Request, conf *configs.Configuration) bool {
	if _, contain := req.Headers["Authorization"]; !contain {
	    return false
	}

	cred, _ := req.Headers["Authorization"];
	// Trim must use conf.Auth.Type.
	if strings.Trim(cred, "Basic ") != "dGVzdDp0ZXN0" {
	    return false
	}

	return true
}
