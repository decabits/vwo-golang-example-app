package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/decabits/vwo-golang-example-app/config"
	"github.com/decabits/vwo-golang-example-app/server"
)

func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	server.Init()
}
