package configs

import (
	"fmt"
	"os"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var _ = fmt.Println // For debugging;
var _ = os.Args // For debugging;

//var currentDir = getCurrentDir()
//func getCurrentDir() string {
//	d, _ := os.Getwd()
//	return d
//}

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

func Load(confFilePath string) *Configuration {
    //fmt.Println(currentDir)
    var conf *Configuration
    buf, err := ioutil.ReadFile(confFilePath)
    if err != nil {
	panic(err)
    }

    err = yaml.UnmarshalStrict(buf, &conf)
    if err != nil {
	panic(err)
    }

    return conf
}
