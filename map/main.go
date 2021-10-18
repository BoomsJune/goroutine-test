package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	m.Store("A", "a")
	m.Store("B", "b")

	val, _ := m.Load("A")
	fmt.Println(val)
}
