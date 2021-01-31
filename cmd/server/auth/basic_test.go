package auth

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestImportUserDb(t *testing.T) {
    fmt.Println("TestimportUserDb")

    importUserDb()
    t.Log(Users)
}

func TestDecodeBase64(t *testing.T) {
    fmt.Println("TestDecodeBase64")

    result := decodeBase64("dGVzdDpwYXNzd29yZA==")
    t.Log(result)
}

func TestAuthCheck(t *testing.T) {
    fmt.Println("TestAuthCheck")

    result, msg := AuthUser(base64.StdEncoding.EncodeToString([]byte("test:password")))
    t.Log(result)
    t.Log(msg)
    result, msg = AuthUser(base64.StdEncoding.EncodeToString([]byte("test:invalidPass")))
    t.Log(result)
    t.Log(msg)
    result, msg = AuthUser(base64.StdEncoding.EncodeToString([]byte("notFoundUser:password")))
    t.Log(result)
    t.Log(msg)
}
