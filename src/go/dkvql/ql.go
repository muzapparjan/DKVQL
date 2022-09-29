package dkvql

type Sentence interface{}

type Insert struct {
	Key     *Token
	Value   *Token
	Timeout *Token
}

type Query struct {
	Prefix  *Token
	Key     *Token
	Timeout *Token
}

type Update struct {
	Prefix  *Token
	Key     *Token
	Value   *Token
	Timeout *Token
}

type Delete struct {
	Prefix  *Token
	Key     *Token
	Timeout *Token
}

type Listen struct {
	Action   *Token
	Prefix   *Token
	Key      *Token
	Callback *Token
}
