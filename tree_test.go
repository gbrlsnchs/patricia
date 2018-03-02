package patricia_test

import (
	"strconv"
	"sync"
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
			expected:     false,
			expectedSize: 1,
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
		// #7
		{
			tree:         New("#7"),
			str:          "testing",
			expected:     false,
			expectedSize: 33,
			handlerFunc: func(t *Tree) {
				t.Add("test", "foo")
				t.Add("testing", "bar")
				t.Del("testing")
			},
		},
		// #8
		{
			tree:          New("#8"),
			str:           "test",
			expected:      true,
			expectedSize:  33,
			expectedValue: "foo",
			handlerFunc: func(t *Tree) {
				t.Add("test", "foo")
				t.Add("testing", "bar")
				t.Del("testing")
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

func TestRace(t *testing.T) {
	list := []string{
		"foo",
		"bar",
		"foobar",
		"foobarbaz",
		"qux",
		"barbazqux",
	}
	tree := New("TestRace")
	tree.Safe = true
	var wg sync.WaitGroup

	wg.Add(len(list))

	for i, n := range list {
		go func(i int, n string) {
			defer wg.Done()
			tree.Add(n, i)

			_ = tree.Get(n)
		}(i, n)
	}

	wg.Wait()
}
