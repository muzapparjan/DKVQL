package dkvql

import (
	"testing"
)

func TestNFA(t *testing.T) {
	n, _ := getNFA()
	t.Log(n.string())
	t.Fail()
}
