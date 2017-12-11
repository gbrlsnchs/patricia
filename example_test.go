package patricia_test

import (
	"fmt"
	"time"

	"github.com/gbrlsnchs/patricia"
)

func Example() {
	// Example from https://upload.wikimedia.org/wikipedia/commons/a/ae/Patricia_trie.svg.
	t := patricia.New("Example")

	t.Add("romane", 1)
	t.Add("romanus", 2)
	t.Add("romulus", 3)
	t.Add("rubens", 4)
	t.Add("ruber", 5)
	t.Add("rubicon", 6)
	t.Add("rubicundus", 7)

	n := t.Get("romanus")

	fmt.Println(n.Value)
	// Output: 2
}

func ExampleTree_Safe() {
	list := []string{
		"romane",
		"romanus",
		"romulus",
		"rubens",
		"ruber",
		"rubicon",
		"rubicundus",
	}
	tree := patricia.New("TestRace")
	tree.Safe = true

	for i, n := range list {
		go func(i int, n string) {
			tree.Add(n, i)
			time.Sleep(time.Second * 3)
		}(i+1, n)
	}

	for _, n := range list {
		go func(n string) {
			_ = tree.Get(n)

			time.Sleep(time.Second * 3)
		}(n)
	}
}
