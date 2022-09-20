package dkvql

import (
	"errors"
	"strings"
)

func addKeyword(nfa *NFA, start string, keyword string, priority float32) error {
	if keyword == "" {
		return errors.New("AddKeyword: keyword must not be empty")
	}
	keyword = strings.ToLower(keyword)
	runes := []rune(keyword)
	length := len(runes)
	var err error
	if length == 1 {
		err = nfa.addState(keyword, priority, true)
		if err != nil {
			return err
		}
		err = nfa.addTransition(runes[0], start, keyword)
		return err
	}
	//TODO
	return nil
}
