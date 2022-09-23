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

	tokens, err := dkvql.Lex("insert Query UPDATE delete")
	if err != nil {
		panic(err)
	}

	for _, token := range tokens {
		fmt.Println(token.String())
	}
}

func parseArgs(args []string) error {
	/*for _, arg := range args {
		//TODO
	}*/
	return nil
}
