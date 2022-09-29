package dkvql

import (
	"fmt"
)

func Lex(src string) ([]*Token, error) {
	n, err := getNFA()
	if err != nil {
		return nil, err
	}
	tokens, err := lex(src, n)
	return tokens, err
}

func lex(src string, n *nfa) ([]*Token, error) {
	input := []rune(src)
	tokens := make([]*Token, 0)
	length := len(input)
	cursor := 0
	line := 1
	index := 1
	var c rune
	output := make([]rune, 0)
	n.reset(initialStates)
	for cursor < length {
		c = input[cursor]
		if len(output) == 0 {
			for s := range skip {
				if c == s {
					cursor++
					index++
					if c == '\n' {
						line++
						index = 0
					}
					continue
				}
			}
		}
		err := n.input(c)
		if err != nil {
			return nil, err
		}
		output = append(output, c)
		if cursor == length-1 {
			if accept, best := n.accept(); accept {
				tokens = append(tokens, newToken(best.name, string(output)))
			}
			break
		}
		cursor++
		index++
		if failed := n.failed(); failed {
			for {
				cursor--
				index--
				err := n.back()
				if err != nil {
					return nil, err
				}
				if len(output) == 1 {
					output = make([]rune, 0)
				} else {
					output = output[:len(output)-1]
				}
				if len(output) == 0 {
					return nil, fmt.Errorf("Lex: failed to parse token at line %v pos %v", line, index)
				}
				if accept, best := n.accept(); accept {
					tokens = append(tokens, newToken(best.name, string(output)))
					output = make([]rune, 0)
					n.reset(initialStates)
					cursor++
					index++
					break
				}
			}
		}
	}
	return tokens, nil
}
