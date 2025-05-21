package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

func main() {
	var cfg string

	// verbose := pflag.Bool("verbose", false, "verbose output")
	verbose := pflag.BoolP("verbose", "v", false, "verbose output")

	// pflag.StringVar(&cfg, "config.yaml", "config file")
	pflag.StringVarP(&cfg, "cfg", "c", "config.yaml", "config file")
	pflag.Lookup("cfg").NoOptDefVal = "noconfig.json"

	pflag.Parse()

	if *verbose {
		fmt.Println("loading config from:", cfg)
	} else {
		fmt.Println(cfg)
	}
}

// go run main.go
// go run main.go -v
// go run main.go --verbose
// go run main.go -v=false
// go run main.go -v=true --cfg=123
