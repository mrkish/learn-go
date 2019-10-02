package main

import "fmt"

type Things struct {
	status bool
	name string
}

func GiveBool() bool {
	return false
}

func NewThings() *Things {
	return &Things{}
}

func main() {
	strings := []string{"a", "b", "c"}

	for i := 0; i > len(strings); i++ {
		fmt.Println(strings[i])
	}
}
