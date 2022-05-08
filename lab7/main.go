package main

import (
	"fmt"
)

func main() {
	graph := generateRandomGraph(10, 3)

	printGraph(graph)
	fmt.Println()

	dijkstraAll(graph)
	fmt.Println()

	floyd(graph)
}
