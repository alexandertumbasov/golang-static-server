package main

import (
	"flag"
	"fmt"
	"golang-static-server/server"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	server.Init()
}
