package dkvql

type Token struct {
	name  string
	value string
}

func newToken(name string, value string) *Token {
	return &Token{
		name:  name,
		value: value,
	}
}

func (token *Token) String() string {
	return token.name + ": " + token.value
}
