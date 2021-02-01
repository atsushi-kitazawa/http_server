package auth

import (
	"strings"

	"github.com/atsushi-kitazawa/http_server/cmd/server/request"
	"github.com/atsushi-kitazawa/http_server/configs"
)

func InitAuthModule() {
    authType := configs.Conf.Auth.Type
    switch authType {
        case "Basic":
	    importUserDb()
	default:
	    panic("fatal error.")
	}
}

// chech request resource is required authorization.
func IsAuthRequireResource(req *request.Request) bool {
   if strings.Contains(req.Resource, configs.Conf.Auth.Location) {
       return true
   }
   return false
}

// chech Authorization header value is valid (support only Basic Auth)
func CheckAuth(req *request.Request) bool {
	if _, contain := req.Headers["Authorization"]; !contain {
	    return false
	}

	cred, _ := req.Headers["Authorization"];
	// Trim must use conf.Auth.Type.
	if ok, _ := AuthUser(strings.Trim(cred, "Basic ")); ok {
	    return true
	}

	return false
}
