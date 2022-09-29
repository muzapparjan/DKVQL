package dkvql

import (
	"fmt"
)

func Parse(src string) (Sentence, error) {
	tokens, err := Lex(src)
	if err != nil {
		return nil, err
	}
	return parse(tokens)
}

func parse(tokens []*Token) (Sentence, error) {
	mayMatch := make(map[rule]struct{})
	for rule := range rules {
		if rule.MayMatch(tokens) {
			mayMatch[rule] = struct{}{}
		}
	}
	for rule := range mayMatch {
		sentence, err := rule.Match(tokens)
		if err != nil {
			continue
		}
		return sentence, nil
	}
	return nil, fmt.Errorf("\nfailed to parse tokens %v", tokens)
}
