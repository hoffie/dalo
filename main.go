package main

import (
	"fmt"
	"os"
)

func main() {
	path := os.Getenv("DALO_DB")
	if path == "" {
		fmt.Printf("error: please set DALO_DB environment variable\n")
		os.Exit(1)
	}
	cli := CLI{DBFile: path}
	os.Exit(cli.Run(os.Args[1:]))
}
