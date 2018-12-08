package main

import (
	"fmt"
	"hound/diffs"
)

func main() {
	excess, missing := diffs.DiffSetSlices([]string{"a", "b"}, []string{"c", "b"})
	fmt.Println(excess)
	fmt.Println(missing)
}
