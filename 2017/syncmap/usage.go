package main

import (
	"fmt"
	"sync"
)

func main() {
	// Initialize
	var m sync.Map

	// Store something
	m.Store("name", "Gopher")

	name, ok := m.Load("name")
	if ok {
		fmt.Printf("Name is %v, type is %T\n", name, name)
	}
}
