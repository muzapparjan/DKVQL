package dkvql

const (
	start string = "__start__"
)

var (
	keywords map[string]float32 = map[string]float32{
		"add": 10,
	}
)

func CreateNFA() (*NFA, error) {
	nfa := newNFA()

	err := nfa.addState(start, 0, false)
	if err != nil {
		return nil, err
	}

	err = nfa.setInitialStates(start)
	if err != nil {
		return nil, err
	}

	for keyword, priority := range keywords {
		err = addKeyword(nfa, start, keyword, priority)
		if err != nil {
			return nil, err
		}
	}
	//TODO
	return nfa, nil
}
