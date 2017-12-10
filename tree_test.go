package patricia_test

import (
	"strconv"
	"testing"

	. "github.com/gbrlsnchs/patricia"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	a := assert.New(t)
	tests := []*struct {
		tree           *Tree
		handlerFunc    func(*Tree)
		str            string
		ph             rune
		delim          rune
		expected       bool
		expectedSize   uint
		expectedValue  interface{}
		expectedParams map[string]string
	}{
		// #0
		{
			tree:         New("#0"),
			str:          "test",
			expectedSize: 1,
		},
		// #1
		{
			tree:         New("#1"),
			str:          "test",
			expected:     true,
			expectedSize: 33,
			handlerFunc: func(t *Tree) {
				t.Add("test", nil)
			},
		},
		// #2
		{
			tree:          New("#2"),
			str:           "test",
			expected:      true,
			expectedSize:  57,
			expectedValue: "foo",
			handlerFunc: func(t *Tree) {
				t.Add("test", "foo")
				t.Add("testing", "bar")
			},
		},
		// #3
		{
			tree:          New("#3"),
			str:           "testing",
			expected:      true,
			expectedSize:  57,
			expectedValue: "bar",
			handlerFunc: func(t *Tree) {
				t.Add("test", "foo")
				t.Add("testing", "bar")
			},
		},
		// #4
		{
			tree:          New("#4"),
			str:           "tester",
			expected:      false,
			expectedSize:  57,
			expectedValue: "bar",
			handlerFunc: func(t *Tree) {
				t.Add("test", "foo")
				t.Add("testing", "bar")
			},
		},
		// #5
		{
			tree:          New("#5"),
			str:           "火",
			expected:      true,
			expectedSize:  25,
			expectedValue: "foo",
			handlerFunc: func(t *Tree) {
				t.Add("火", "foo")
			},
		},
		// #6
		{
			tree:          New("#6"),
			str:           "火",
			expected:      true,
			expectedSize:  42,
			expectedValue: "foo",
			handlerFunc: func(t *Tree) {
				t.Add("火", "foo")
				t.Add("水", "bar")
			},
		},
	}

	for i, test := range tests {
		index := strconv.Itoa(i)

		if test.handlerFunc != nil {
			test.handlerFunc(test.tree)
		}

		a.Exactly(test.expectedSize, test.tree.Size(), index)

		n := test.tree.Get(test.str)

		a.Exactly(test.expected, n != nil, index)

		if n != nil {
			a.Exactly(test.expectedValue, n.Value, index)
			t.Logf("n.Value = %#v\n", n.Value)
		}
	}
}
