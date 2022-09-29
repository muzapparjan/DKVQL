package dkvql

const (
	nfaStart   string = "__start__"
	nfaEpsilon rune   = -1
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
	binaryNumbers map[rune]struct{} = map[rune]struct{}{
		'0': {},
		'1': {},
	}
	decimalNumbers map[rune]struct{} = map[rune]struct{}{
		'0': {},
		'1': {},
		'2': {},
		'3': {},
		'4': {},
		'5': {},
		'6': {},
		'7': {},
		'8': {},
		'9': {},
	}
	hexNumbers map[rune]struct{} = map[rune]struct{}{
		'0': {},
		'1': {},
		'2': {},
		'3': {},
		'4': {},
		'5': {},
		'6': {},
		'7': {},
		'8': {},
		'9': {},
		'a': {},
		'b': {},
		'c': {},
		'd': {},
		'e': {},
		'f': {},
		'A': {},
		'B': {},
		'C': {},
		'D': {},
		'E': {},
		'F': {},
	}
	initialStates map[string]struct{} = map[string]struct{}{
		nfaStart: {},
	}
)
