package configs

import (
    "testing"
    "flag"
)

func TestLoad(t *testing.T) {
    flag.CommandLine.Set("conf", "configuration.yaml")
    flag.Parse()

    var conf Configuration
    Load(&conf)

    t.Log(conf.Ip)
    t.Log(conf.Port)
    t.Log(conf.Auth.Type)
    t.Log(conf.Auth.File)
    t.Log(conf.Auth.Location)
}
