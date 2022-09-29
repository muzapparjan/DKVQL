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

	tokens, err := dkvql.Lex("123456    \n987.654 Delete UPDATE 0b10011101 0xA34b7901 @hello @_hi @how_are_u_123_hi \"Hello!\" 123 \"Stupid Dunkey~\"")
	if err != nil {
		panic(err)
	}

	for _, token := range tokens {
		fmt.Print("\n" + token.String())
	}
}

func parseArgs(args []string) error {
	/*for _, arg := range args {
		//TODO
	}*/
	return nil
}
