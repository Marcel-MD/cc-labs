package main

import (
	"fmt"
	"time"
)

func main() {

	var start time.Time
	var duration time.Duration

	for n := 20; n <= 400; n += 20 {
		graph := generateRandomGraph(n, 3)

		start = time.Now()
		dijkstraAll(graph)
		duration = time.Since(start)
		fmt.Printf("Dijkstra for %d vertices took %d μs to compute\n", n, duration.Microseconds())

		start = time.Now()
		floyd(graph)
		duration = time.Since(start)
		fmt.Printf("Floyd for %d vertices took %d μs to compute\n\n", n, duration.Microseconds())
	}
}
