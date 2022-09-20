package dkvql

import (
	"fmt"
)

type nfaState struct {
	name       string
	priority   float32
	acceptable bool
}

type NFA struct {
	states        map[string]nfaState
	initialStates map[string]struct{}
	transitions   map[string]map[rune]string

	currentStates map[string]struct{}
	history       []map[string]struct{}
}

func newNFAState(name string, priority float32, acceptable bool) nfaState {
	return nfaState{
		name:       name,
		priority:   priority,
		acceptable: acceptable,
	}
}

func newNFA() *NFA {
	return &NFA{
		states:        make(map[string]nfaState),
		initialStates: make(map[string]struct{}),
		transitions:   make(map[string]map[rune]string),
		currentStates: make(map[string]struct{}),
		history:       make([]map[string]struct{}, 0),
	}
}

func (nfa *NFA) addState(name string, priority float32, acceptable bool) error {
	if _, exist := nfa.states[name]; !exist {
		state := newNFAState(name, priority, acceptable)
		nfa.states[name] = state
		return nil
	}
	return fmt.Errorf("NFA.AddState: State named %v already exist", name)
}

func (nfa *NFA) addTransition(input rune, from string, to string) error {
	var (
		set   map[rune]string
		exist bool
	)
	if set, exist = nfa.transitions[from]; !exist {
		set = make(map[rune]string)
		set[input] = to
		nfa.transitions[from] = set
		return nil
	}
	if _, exist = set[input]; !exist {
		set[input] = to
		return nil
	}
	return fmt.Errorf("NFA.AddTransition: Transition {input: %v, from: %v, to: %v} already exist", input, from, to)
}

func (nfa *NFA) setInitialStates(states ...string) error {
	initialStates := make(map[string]struct{})
	for _, state := range states {
		if _, exist := nfa.states[state]; !exist {
			return fmt.Errorf("SetInitialStates: State named %v does not exist.", state)
		}
		initialStates[state] = struct{}{}
	}
	nfa.initialStates = initialStates
	return nil
}

func (nfa *NFA) reset() error {
	nfa.currentStates = make(map[string]struct{})
	nfa.history = make([]map[string]struct{}, 0)
	return nil
}

func (nfa *NFA) input(c rune) error {
	if len(nfa.currentStates) == 0 {
		return nil
	}
	states := make(map[string]struct{})
	for state := range nfa.currentStates {
		if transition, exist := nfa.transitions[state]; exist {
			if to, exist := transition[c]; exist {
				if _, exist := nfa.states[to]; !exist {
					return fmt.Errorf("NFA.Input: failed to transition from state %v to unknown state %v", state, to)
				}
				states[to] = struct{}{}
			}
		}
	}
	nfa.history = append(nfa.history, nfa.currentStates)
	nfa.currentStates = states
	return nil
}

func (nfa *NFA) back() error {
	if len(nfa.history) == 0 {
		return nil
	}
	last := len(nfa.history) - 1
	nfa.currentStates = nfa.history[last]
	nfa.history = nfa.history[:last]
	return nil
}

func (nfa *NFA) accept() bool {
	for state := range nfa.currentStates {
		if s, exist := nfa.states[state]; exist {
			if s.acceptable {
				return true
			}
		}
	}
	return false
}
