package dkvql

import (
	"fmt"
	"math"
)

type nfaState struct {
	name       string
	priority   float32
	acceptable bool
}

type nfa struct {
	states      map[string]nfaState
	transitions map[string]map[rune]map[string]struct{}

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

func newNFA() *nfa {
	return &nfa{
		states:        make(map[string]nfaState),
		transitions:   make(map[string]map[rune]map[string]struct{}),
		currentStates: make(map[string]struct{}),
		history:       make([]map[string]struct{}, 0),
	}
}

func (n *nfa) addState(name string, priority float32, acceptable bool) {
	n.states[name] = newNFAState(name, priority, acceptable)
}

func (n *nfa) addTransition(input rune, from string, to string) {
	var (
		set   map[rune]map[string]struct{}
		tos   map[string]struct{}
		exist bool
	)
	if set, exist = n.transitions[from]; !exist {
		set = make(map[rune]map[string]struct{})
		n.transitions[from] = set
	}
	if tos, exist = set[input]; !exist {
		tos = make(map[string]struct{})
		set[input] = tos
	}
	tos[to] = struct{}{}
}

func (n *nfa) reset(initialStates map[string]struct{}) error {
	n.currentStates = initialStates
	n.history = make([]map[string]struct{}, 0)
	return nil
}

func (n *nfa) input(c rune) error {
	if len(n.currentStates) == 0 {
		return nil
	}
	states := make(map[string]struct{})
	for state := range n.currentStates {
		if transition, exist := n.transitions[state]; exist {
			if tos, exist := transition[c]; exist {
				for to := range tos {
					if _, exist := n.states[to]; !exist {
						return fmt.Errorf("NFA.Input: failed to transition from state %v to unknown state %v", state, to)
					}
					states[to] = struct{}{}
				}
			}
			if tos, exist := transition[nfaEpsilon]; exist {
				for to := range tos {
					states[to] = struct{}{}
				}
			}
		}
	}
	n.history = append(n.history, n.currentStates)
	n.currentStates = states
	return nil
}

func (n *nfa) back() error {
	if len(n.history) == 0 {
		return nil
	}
	last := len(n.history) - 1
	n.currentStates = n.history[last]
	n.history = n.history[:last]
	return nil
}

func (n *nfa) accept() (bool, nfaState) {
	accepted := false
	best := nfaState{
		name:       "NotAccepted",
		priority:   -math.MaxFloat32,
		acceptable: false,
	}
	for state := range n.currentStates {
		if s, exist := n.states[state]; exist {
			if s.acceptable {
				accepted = true
				if best.priority < s.priority {
					best = s
				}
			}
		}
	}
	return accepted, best
}

func (n *nfa) failed(terminateStates map[string]struct{}) bool {
	if accepted, state := n.accept(); accepted {
		if _, exist := terminateStates[state.name]; exist {
			return true
		}
	}
	return len(n.currentStates) == 0
}
