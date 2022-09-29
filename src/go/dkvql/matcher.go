package dkvql

type matcherFunc func([]*Token, *int) error

type matcher struct {
	funcs []matcherFunc
}

func newMatcher() *matcher {
	return &matcher{
		funcs: make([]matcherFunc, 0),
	}
}

func (m *matcher) append(fn matcherFunc) {
	m.funcs = append(m.funcs, fn)
}

func (m *matcher) match(tokens []*Token) error {
	cursor := 0
	for _, fn := range m.funcs {
		err := fn(tokens, &cursor)
		if err != nil {
			return err
		}
	}
	return nil
}
