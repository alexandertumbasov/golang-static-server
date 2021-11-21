package main

import (
	"flag"
	"fmt"
	"golang-static-server/src/server"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: src -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	server.Init()
}
