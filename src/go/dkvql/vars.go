package dkvql

const (
	nfaStart string = "__start__"
)

var (
	keywords map[string]float32 = map[string]float32{
		"insert":   10,
		"query":    10,
		"update":   10,
		"delete":   10,
		"value":    10,
		"prefix":   10,
		"timeout":  10,
		"callback": 10,
		"listen":   10,
	}
	skip map[rune]struct{} = map[rune]struct{}{
		' ':  {},
		'\r': {},
		'\n': {},
	}
	initialStates map[string]struct{} = map[string]struct{}{
		nfaStart: {},
	}
)
