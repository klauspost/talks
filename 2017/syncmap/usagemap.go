package main

import (
	"fmt"
)

func main() {
	// Initialize
	var m = make(map[string]string)

	// Store something
	m["name"] = "Gopher"

	name, ok := m["name"]
	if ok {
		fmt.Printf("Name is %v, type is %T (of course)\n", name, name)
	}
}
