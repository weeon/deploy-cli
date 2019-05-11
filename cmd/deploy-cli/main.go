package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/weeon/deploy-cli/internal/biz"
)

var (
	action string
)

func main() {
	var err error
	err = biz.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	flag.StringVar(&action, "a", "", "action")
	flag.Parse()

	switch action {
	case "deploy":
		biz.Deploy()
	default:
		fmt.Println("no action input")
	}

}
