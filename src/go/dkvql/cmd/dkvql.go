package main

import (
	"dkvql"
	"fmt"
	"os"
)

func main() {
	err := parseArgs(os.Args[1:])
	if err != nil {
		panic(err)
	}

	src := "insert \"username\" value \"admin\" timeout 30"

	sentence, err := dkvql.Parse(src)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nSentence: {%v}", sentence)
}

func parseArgs(args []string) error {
	/*for _, arg := range args {
		//TODO
	}*/
	return nil
}
