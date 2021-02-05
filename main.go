package main

import (
	"flag"
	"fmt"
	"gochat/connector"
)

func main() {
	// load server module
	var m string
	flag.StringVar(&m, "m", "", "run module")
	flag.Parse()
	if m == "" {
		m = "all"
	}
	switch m {
	case "tcp":
		connector.TcpRun()
	default:
		fmt.Println("fail: module param error")
		return
	}
	fmt.Println("staring: run module " + m)
	for {

	}
}
