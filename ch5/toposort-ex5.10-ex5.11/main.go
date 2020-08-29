package main

import (
	"fmt"
	"log"
)

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d\t %s\n", i+1, course)
	}
}

var prereqs = map[string]map[string]bool{
	"algorithms": {"data structures": true},
	"calculus":   {"linear algebra": true},
	"compilers": {
		"data structures":        true,
		"formal languages":       true,
		"computer organizations": true,
	},
	"data structures":  {"discrete math": true},
	"database":         {"data structures": true},
	"discrete math":    {"intro to programming": true},
	"formal languages": {"discrete math": true},
	"networks":         {"operating systems": true},
	"operating systems": {
		"data structures":        true,
		"computer organizations": true,
	},
	"programming languages": {
		"data structures":        true,
		"computer organizations": true,
	},
	"linear algebra": {"basic math": true},
	"basic math":     {"calculus": true},
}

func topoSort(m map[string]map[string]bool) []string {
	var order []string
	// using stash to mark if the item were seen and put on the order list,
	// false means seen but not on the order list,
	// true means already on the order list
	var stash = make(map[string]bool)
	var visitAll func(map[string]bool)
	visitAll = func(items map[string]bool) {
		for item, _ := range items {
			v, ok := stash[item]
			if !ok {
				stash[item] = false
				visitAll(m[item])
				// mark as on list
				stash[item] = true
				order = append(order, item)
			} else {
				if !v { // item was seen but not on list, means cycle
					log.Fatalf("detected cycle prerequisites,"+
						" %s is involved\n",
						item)
				}
			}
		}

	}

	var keys = make(map[string]bool)
	for k, _ := range m {
		keys[k] = true
	}
	visitAll(keys)
	return order
}
