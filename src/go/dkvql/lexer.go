package dkvql

import "fmt"

func Lex(src string, nfa *NFA) ([]*Token, error) {
	input := []rune(src)
	tokens := make([]*Token, 0)
	length := len(input)
	cursor := 0
	var c rune
	for cursor < length {
		c = input[cursor]
		fmt.Print(c)
		//TODO
	}
	return tokens, nil
}
