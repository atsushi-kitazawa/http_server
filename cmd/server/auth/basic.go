package auth

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

var users []UserDb = make([]UserDb, 0) 

type UserDb struct {
    user string
    pass string
}

func init() {
    importUserDb()
}

func decodeBase64(src string) string {
    dec, err := base64.StdEncoding.DecodeString(src)
    if err != nil {
	panic(err)
    }
    return string(dec)
}

func importUserDb() {
    file, err :=  os.Open("user_db")
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
    //fmt.Println(users)
}

func AuthUser(cred string) (bool, string) {
    dec := decodeBase64(cred)
    user := strings.Split(dec, ":")[0]
    pass := strings.Split(dec, ":")[1]
    for _, val := range users {
	if user == val.user {
	    if pass == val.pass {
		return true, "Auth OK"
	    } else {
		return false, fmt.Sprintf("%s user password invalid", user)
	    }
	}
    }
    fmt.Printf("%s user not foud.", user)
    return false, fmt.Sprintf("%s user not found", user)
}
