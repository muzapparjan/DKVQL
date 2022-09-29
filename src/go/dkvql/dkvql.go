package dkvql

var defaultNFA *nfa

func getNFA() (*nfa, error) {
	if defaultNFA == nil {
		n, err := buildNFA()
		if err != nil {
			return nil, err
		}
		defaultNFA = n
	}
	return defaultNFA, nil
}

func buildNFA() (*nfa, error) {
	n := newNFA()
	n.addState(nfaStart, -1, false)
	for keyword, priority := range keywords {
		n.addKeyword(nfaStart, keyword, priority)
	}
	n.addBinaryNumber(5)
	n.addHexNumber(5)
	n.addFPNumber(5)
	n.addDecimalNumber(5)
	//TODO
	return n, nil
}
