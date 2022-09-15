package main

import (
	"os"
)

func main() {
	err := parseArgs(os.Args[1:])
	if err != nil {
		panic(err)
	}
}

func parseArgs(args []string) error {
	for _, arg := range args {
		//TODO
	}
	return nil
}
