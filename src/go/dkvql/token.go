package dkvql

type Token struct {
	Name  string
	Value string
}

func newToken(name string, value string) *Token {
	return &Token{
		Name:  name,
		Value: value,
	}
}

func (token *Token) String() string {
	return token.Name + ": " + token.Value
}
