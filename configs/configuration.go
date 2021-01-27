package configs

import (
	"fmt"
	"os"
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"github.com/atsushi-kitazawa/http_server/cmd/server/enviroment"
)

var currentDir = getCurrentDir()

func getCurrentDir() string {
	d, _ := os.Getwd()
	return d
}
type Configuration struct {
    Ip string `yaml:"ip"`
    Port string `yaml:"port"`
    Auth Authentication `yaml:"auth"`
}

type Authentication struct {
    Type string `yaml:"type"`
    File string `yaml:"file"`
    Location string `yaml:"location"`
}

func Load(conf *Configuration) {
    fmt.Println(currentDir)
    buf, err := ioutil.ReadFile(enviroment.GetEnv().ConfFile)
    if err != nil {
	panic(err)
    }

    err = yaml.UnmarshalStrict(buf, &conf)
    if err != nil {
	panic(err)
    }
    //fmt.Println("ip>", conf.Ip)
    //fmt.Println("port>", conf.Port)
}
