package configs

import (
	"testing"
)

func TestLoad(t *testing.T) {
    var conf *Configuration
    conf = Load("configuration.yaml")

    t.Log(conf.Ip)
    t.Log(conf.Port)
    t.Log(conf.Auth.Type)
    t.Log(conf.Auth.File)
    t.Log(conf.Auth.Location)
}
