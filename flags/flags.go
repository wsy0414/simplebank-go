package flags

import (
	"flag"
	"fmt"
	"testing"
)

var (
	ENV string
)

func init() {
	flag.StringVar(&ENV, "env", "", "to set env")
	testing.Init()
	flag.Parse()
	fmt.Println(ENV)
}
