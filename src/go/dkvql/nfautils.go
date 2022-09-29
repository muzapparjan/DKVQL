package dkvql

import (
	"strings"
)

func (n *nfa) addKeyword(start string, keyword string, priority float32) {
	if keyword == "" {
		return
	}
	lower := []rune(strings.ToLower(keyword))
	upper := []rune(strings.ToUpper(keyword))
	if len(lower) == 1 {
		n.addState(keyword, priority, true)
		n.addTransition(lower[0], start, keyword)
		n.addTransition(upper[0], start, keyword)
		return
	}
	prefix := "__" + keyword + "__"
	if len(lower) == 2 {
		n.addState(keyword, priority, true)
		lowerState := prefix + string(lower[0])
		upperState := prefix + string(upper[0])
		n.addState(lowerState, -1, false)
		n.addState(upperState, -1, false)
		n.addTransition(lower[0], start, lowerState)
		n.addTransition(upper[0], start, upperState)
		n.addTransition(lower[1], lowerState, keyword)
		n.addTransition(upper[1], upperState, keyword)
		n.addTransition(lower[1], upperState, keyword)
		return
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
		n.addState(lowerChain[i], -1, false)
		n.addTransition(lower[i-1], lowerChain[i-1], lowerChain[i])
		n.addState(upperChain[i], -1, false)
		n.addTransition(upper[i-1], upperChain[i-1], upperChain[i])
	}
	n.addState(keyword, priority, true)
	n.addTransition(lower[len(lower)-1], lowerChain[len(lowerChain)-2], keyword)
	n.addTransition(upper[len(upper)-1], upperChain[len(upperChain)-2], keyword)
	n.addTransition(lower[1], upperChain[1], lowerChain[2])
}

func (n *nfa) addBinaryNumber(priority float32) {
	n.addState("binary_number", priority, true)
	n.addState("__bin__1", -1, false)
	n.addState("__bin__2", -1, false)
	n.addTransition('0', nfaStart, "__bin__1")
	n.addTransition('b', "__bin__1", "__bin__2")
	n.addTransition('B', "__bin__1", "__bin__2")
	for r := range binaryNumbers {
		n.addTransition(r, "__bin__2", "binary_number")
		n.addTransition(r, "binary_number", "binary_number")
	}
}

func (n *nfa) addDecimalNumber(priority float32) {
	n.addState("decimal_number", priority, true)
	for r := range decimalNumbers {
		n.addTransition(r, nfaStart, "decimal_number")
		n.addTransition(r, "decimal_number", "decimal_number")
	}
}

func (n *nfa) addHexNumber(priority float32) {
	n.addState("hex_number", priority, true)
	n.addState("__hex__1", -1, false)
	n.addState("__hex__2", -1, false)
	n.addTransition('0', nfaStart, "__hex__1")
	n.addTransition('x', "__hex__1", "__hex__2")
	n.addTransition('X', "__hex__1", "__hex__2")
	for r := range hexNumbers {
		n.addTransition(r, "__hex__2", "hex_number")
		n.addTransition(r, "hex_number", "hex_number")
	}
}

func (n *nfa) addFPNumber(priority float32) {
	n.addState("fp_number", priority, true)
	n.addState("__fp__1", -1, false)
	n.addState("__fp__2", -1, false)
	for r := range decimalNumbers {
		n.addTransition(r, nfaStart, "__fp__1")
		n.addTransition(r, "__fp__1", "__fp__1")
	}
	n.addTransition('.', "__fp__1", "__fp__2")
	for r := range decimalNumbers {
		n.addTransition(r, "__fp__2", "fp_number")
		n.addTransition(r, "fp_number", "fp_number")
	}
}
