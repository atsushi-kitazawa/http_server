package enviroment

import (
    "flag"
)

var env *Enviroment = new(Enviroment)

type Enviroment struct {
    ConfFile string
}

func init() {
    flag.StringVar(&env.ConfFile, "conf", "configs/configuration.yaml", "configuration file")
}

func GetEnv() *Enviroment {
    return env
}
