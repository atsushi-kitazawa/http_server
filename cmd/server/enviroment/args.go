package enviroment

import (
	"flag"
	"fmt"
)

var _ = fmt.Println //For debugging.

var args *Args = new(Args)

type Args struct {
    ConfFile string
}

func init() {
    fmt.Println("args init() call")
    flag.StringVar(&args.ConfFile, "conf", "configs/configuration.yaml", "configuration file")
    flag.Parse()
}

func GetArgs() *Args {
    fmt.Println("enviroment GetArgs() call")
    return args
}
