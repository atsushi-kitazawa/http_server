package auth

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/atsushi-kitazawa/http_server/configs"
)

var users []UserDb = make([]UserDb, 0)

type UserDb struct {
    user string
    pass string
}

func decodeBase64(src string) string {
    dec, err := base64.StdEncoding.DecodeString(src)
    if err != nil {
	panic(err)
    }
    fmt.Println("basic.go debug>", string(dec))
    return string(dec)
}

func importUserDb() {
    conf := configs.Conf
    file, err :=  os.Open(conf.Auth.File)
    if err != nil {
	panic(err)
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    for i := 0 ; ; i++ {
	line, _, err := reader.ReadLine()
	if err == io.EOF {
	    fmt.Println("EOF")
	    break
	} else if err != nil {
	    panic(err)
	}
	userPass := strings.Split(string(line), ":")
	ud := UserDb{userPass[0], userPass[1]}
	users = append(users, ud)
    }
    fmt.Println("basic.go debug>>>", users)
}

func AuthUser(cred string) (bool, string) {
    dec := decodeBase64(cred)
    user := strings.Split(dec, ":")[0]
    pass := strings.Split(dec, ":")[1]
    fmt.Println("basic.go debug>>", user)
    fmt.Println("basic.go debug>>", pass)
    for _, val := range users {
	if user == val.user {
	    if pass == val.pass {
		fmt.Println("auth ok")
		return true, "Auth OK"
	    } else {
		fmt.Printf("%s user password invalid.\n", user)
		return false, fmt.Sprintf("%s user password invalid\n", user)
	    }
	}
    }
    fmt.Printf("%s user not foud.\n", user)
    return false, fmt.Sprintf("%s user not found\n", user)
}
