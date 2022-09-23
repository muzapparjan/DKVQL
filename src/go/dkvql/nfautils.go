package dkvql

import (
	"errors"
	"strings"
)

func (n *nfa) addKeyword(start string, keyword string, priority float32) error {
	if keyword == "" {
		return errors.New("AddKeyword: keyword must not be empty")
	}
	lower := []rune(strings.ToLower(keyword))
	upper := []rune(strings.ToUpper(keyword))
	if len(lower) == 1 {
		err := n.addState(keyword, priority, true)
		if err != nil {
			return err
		}
		err = n.addTransition(lower[0], start, keyword)
		if err != nil {
			return err
		}
		err = n.addTransition(upper[0], start, keyword)
		return err
	}
	prefix := "__" + keyword + "__"
	if len(lower) == 2 {
		err := n.addState(keyword, priority, true)
		if err != nil {
			return err
		}
		lowerState := prefix + string(lower[0])
		upperState := prefix + string(upper[0])
		err = n.addState(lowerState, -1, false)
		if err != nil {
			return err
		}
		err = n.addState(upperState, -1, false)
		if err != nil {
			return err
		}
		err = n.addTransition(lower[0], start, lowerState)
		if err != nil {
			return err
		}
		err = n.addTransition(upper[0], start, upperState)
		if err != nil {
			return err
		}
		err = n.addTransition(lower[1], lowerState, keyword)
		if err != nil {
			return err
		}
		err = n.addTransition(upper[1], upperState, keyword)
		if err != nil {
			return err
		}
		err = n.addTransition(lower[1], upperState, keyword)
		return err
	}
	chainFn := func(input []rune) []string {
		result := make([]string, 0, len(input)+1)
		result = append(result, start)
		for i := range input {
			result = append(result, prefix+string(input[:i+1]))
		}
		return result
	}
	lowerChain := chainFn(lower)
	upperChain := chainFn(upper)
	length := len(lowerChain)
	for i := 1; i < length-1; i++ {
		err := n.addState(lowerChain[i], -1, false)
		if err != nil {
			return err
		}
		err = n.addTransition(lower[i-1], lowerChain[i-1], lowerChain[i])
		if err != nil {
			return err
		}
		err = n.addState(upperChain[i], -1, false)
		if err != nil {
			return err
		}
		err = n.addTransition(upper[i-1], upperChain[i-1], upperChain[i])
		if err != nil {
			return err
		}
	}
	err := n.addState(keyword, priority, true)
	if err != nil {
		return err
	}
	err = n.addTransition(lower[len(lower)-1], lowerChain[len(lowerChain)-2], keyword)
	if err != nil {
		return err
	}
	err = n.addTransition(upper[len(upper)-1], upperChain[len(upperChain)-2], keyword)
	if err != nil {
		return err
	}
	err = n.addTransition(lower[1], upperChain[1], lowerChain[2])
	if err != nil {
		return err
	}
	return nil
}
