package dkvql

import "errors"

type rule interface {
	MayMatch([]*Token) bool
	Match([]*Token) (Sentence, error)
}

type insertRule struct{}

func (rule *insertRule) MayMatch(tokens []*Token) bool {
	if len(tokens) == 0 {
		return false
	}
	return tokens[0].Name == "insert"
}

func (rule *insertRule) Match(tokens []*Token) (Sentence, error) {
	insert := &Insert{}
	m := newMatcher()
	err := errors.New("failed to parse Insert sentence")
	m.append(func(tokens []*Token, cursor *int) error {
		if *cursor >= len(tokens) {
			return err
		}
		if tokens[*cursor].Name != "insert" {
			return err
		}
		*cursor++
		return nil
	})
	m.append(func(tokens []*Token, cursor *int) error {
		if *cursor >= len(tokens) {
			return err
		}
		if _, exist := keyTypes[tokens[*cursor].Name]; !exist {
			return err
		}
		insert.Key = tokens[*cursor]
		*cursor++
		return nil
	})
	m.append(func(tokens []*Token, cursor *int) error {
		if *cursor >= len(tokens)-1 {
			return err
		}
		if tokens[*cursor].Name != "value" {
			return err
		}
		*cursor++
		if _, exist := valueTypes[tokens[*cursor].Name]; !exist {
			return err
		}
		insert.Value = tokens[*cursor]
		*cursor++
		return nil
	})
	m.append(func(tokens []*Token, cursor *int) error {
		if *cursor >= len(tokens) {
			return nil
		}
		if *cursor == len(tokens)-1 {
			return err
		}
		if tokens[*cursor].Name != "timeout" {
			return err
		}
		*cursor++
		if !isNumber(tokens[*cursor]) && tokens[*cursor].Name != "name" {
			return err
		}
		insert.Timeout = tokens[*cursor]
		*cursor++
		return nil
	})
	err = m.match(tokens)
	if err != nil {
		return nil, err
	}
	return *insert, nil
}

type queryRule struct{}

func (rule *queryRule) MayMatch(tokens []*Token) bool {
	if len(tokens) == 0 {
		return false
	}
	return tokens[0].Name == "query"
}

func (rule *queryRule) Match(tokens []*Token) (Sentence, error) {
	query := &Query{}
	m := newMatcher()
	err := errors.New("failed to parse Query sentence")
	m.append(func(tokens []*Token, cursor *int) error {
		if *cursor >= len(tokens) {
			return err
		}
		if tokens[*cursor].Name != "query" {
			return err
		}
		*cursor++
		return nil
	})
	m.append(func(tokens []*Token, cursor *int) error {
		if *cursor >= len(tokens) {
			return err
		}
		if tokens[*cursor].Name == "prefix" {
			query.Prefix = tokens[*cursor]
			*cursor++
		}
		return nil
	})
	m.append(func(tokens []*Token, cursor *int) error {
		if *cursor >= len(tokens) {
			return err
		}
		if _, exist := keyTypes[tokens[*cursor].Name]; !exist {
			return err
		}
		query.Key = tokens[*cursor]
		*cursor++
		return nil
	})
	m.append(func(tokens []*Token, cursor *int) error {
		if *cursor >= len(tokens) {
			return nil
		}
		if *cursor == len(tokens)-1 {
			return err
		}
		if tokens[*cursor].Name != "timeout" {
			return err
		}
		*cursor++
		if !isNumber(tokens[*cursor]) && tokens[*cursor].Name != "name" {
			return err
		}
		query.Timeout = tokens[*cursor]
		*cursor++
		return nil
	})
	err = m.match(tokens)
	if err != nil {
		return nil, err
	}
	return query, nil
}

type updateRule struct{}

func (rule *updateRule) MayMatch(tokens []*Token) bool {
	if len(tokens) == 0 {
		return false
	}
	return tokens[0].Name == "update"
}

func (rule *updateRule) Match(tokens []*Token) (Sentence, error) {
	update := &Update{}
	m := newMatcher()
	err := errors.New("failed to parse Update sentence")
	m.append(func(tokens []*Token, cursor *int) error {
		if *cursor >= len(tokens) {
			return err
		}
		if tokens[*cursor].Name != "update" {
			return err
		}
		*cursor++
		return nil
	})
	m.append(func(tokens []*Token, cursor *int) error {
		if *cursor >= len(tokens) {
			return err
		}
		if tokens[*cursor].Name == "prefix" {
			update.Prefix = tokens[*cursor]
			*cursor++
		}
		return nil
	})
	m.append(func(tokens []*Token, cursor *int) error {
		if *cursor >= len(tokens) {
			return err
		}
		if _, exist := keyTypes[tokens[*cursor].Name]; !exist {
			return err
		}
		update.Key = tokens[*cursor]
		*cursor++
		return nil
	})
	m.append(func(tokens []*Token, cursor *int) error {
		if *cursor >= len(tokens)-1 {
			return err
		}
		if tokens[*cursor].Name != "value" {
			return err
		}
		*cursor++
		if _, exist := valueTypes[tokens[*cursor].Name]; !exist {
			return err
		}
		update.Value = tokens[*cursor]
		*cursor++
		return nil
	})
	m.append(func(tokens []*Token, cursor *int) error {
		if *cursor >= len(tokens) {
			return nil
		}
		if *cursor == len(tokens)-1 {
			return err
		}
		if tokens[*cursor].Name != "timeout" {
			return err
		}
		*cursor++
		if !isNumber(tokens[*cursor]) && tokens[*cursor].Name != "name" {
			return err
		}
		update.Timeout = tokens[*cursor]
		*cursor++
		return nil
	})
	err = m.match(tokens)
	if err != nil {
		return nil, err
	}
	return update, nil
}

type deleteRule struct{}

func (rule *deleteRule) MayMatch(tokens []*Token) bool {
	if len(tokens) == 0 {
		return false
	}
	return tokens[0].Name == "delete"
}

func (rule *deleteRule) Match(tokens []*Token) (Sentence, error) {
	delete := &Delete{}
	m := newMatcher()
	err := errors.New("failed to parse Delete sentence")
	m.append(func(tokens []*Token, cursor *int) error {
		if *cursor >= len(tokens) {
			return err
		}
		if tokens[*cursor].Name != "delete" {
			return err
		}
		*cursor++
		return nil
	})
	m.append(func(tokens []*Token, cursor *int) error {
		if *cursor >= len(tokens) {
			return err
		}
		if tokens[*cursor].Name == "prefix" {
			delete.Prefix = tokens[*cursor]
			*cursor++
		}
		return nil
	})
	m.append(func(tokens []*Token, cursor *int) error {
		if *cursor >= len(tokens) {
			return err
		}
		if _, exist := keyTypes[tokens[*cursor].Name]; !exist {
			return err
		}
		delete.Key = tokens[*cursor]
		*cursor++
		return nil
	})
	m.append(func(tokens []*Token, cursor *int) error {
		if *cursor >= len(tokens) {
			return nil
		}
		if *cursor == len(tokens)-1 {
			return err
		}
		if tokens[*cursor].Name != "timeout" {
			return err
		}
		*cursor++
		if !isNumber(tokens[*cursor]) && tokens[*cursor].Name != "name" {
			return err
		}
		delete.Timeout = tokens[*cursor]
		*cursor++
		return nil
	})
	err = m.match(tokens)
	if err != nil {
		return nil, err
	}
	return delete, nil
}

type listenRule struct{}

func (rule *listenRule) MayMatch(tokens []*Token) bool {
	if len(tokens) == 0 {
		return false
	}
	return tokens[0].Name == "listen"
}

func (rule *listenRule) Match(tokens []*Token) (Sentence, error) {
	//TODO
	return nil, nil
}

func isNumber(token *Token) bool {
	switch token.Name {
	case "binary_number":
		return true
	case "decimal_number":
		return true
	case "hex_number":
		return true
	case "fp_number":
		return true
	default:
		return false
	}
}
