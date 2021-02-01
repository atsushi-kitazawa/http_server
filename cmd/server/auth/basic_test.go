package auth

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"

	"github.com/atsushi-kitazawa/http_server/configs"
	"github.com/stretchr/testify/assert"
)

var _ = os.Args // For debuggin.

var conf *configs.Configuration = configs.Load("../../../configs/configuration.yaml")

//func TestMain(m *testing.M) {
//
//    code := m.Run()
//
//    os.Exit(code)
//}

func TestImportUserDb(t *testing.T) {
    fmt.Println("TestimportUserDb")

    importUserDb()

    assert.Equal(t, users[0].user, "test")
    assert.Equal(t, users[0].pass, "password")
    assert.Equal(t, users[1].user, "user1")
    assert.Equal(t, users[1].pass, "pass")
}

func TestDecodeBase64(t *testing.T) {
    fmt.Println("TestDecodeBase64")

    result := decodeBase64("dGVzdDpwYXNzd29yZA==")
    assert.Equal(t, result, "test:password")
}

func TestAuthCheck(t *testing.T) {
    fmt.Println("TestAuthCheck")

    result, _ := AuthUser(base64.StdEncoding.EncodeToString([]byte("test:password")))
    assert.Equal(t, result, true)
    result, _ = AuthUser(base64.StdEncoding.EncodeToString([]byte("test:invalidPass")))
    assert.Equal(t, result, false)
    result, _ = AuthUser(base64.StdEncoding.EncodeToString([]byte("notFoundUser:password")))
    assert.Equal(t, result, false)
}
