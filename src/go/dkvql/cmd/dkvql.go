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

	nfa, err := dkvql.CreateNFA()
	if err != nil {
		panic(err)
	}

	tokens, err := dkvql.Lex("", nfa)
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
