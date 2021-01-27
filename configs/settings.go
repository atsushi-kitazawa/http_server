package configs

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Settings struct {
    Ip string `yaml:"ip"`
    Port string `yaml:"port"`
}

func Load(settings *Settings) {
    buf, err := ioutil.ReadFile("configs/settings.yaml")
    if err != nil {
	panic(err)
    }

    err = yaml.UnmarshalStrict(buf, &settings)
    if err != nil {
	panic(err)
    }
    //fmt.Println("ip>", settings.Ip)
    //fmt.Println("port>", settings.Port)
}
