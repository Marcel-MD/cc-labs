package main

import (
	"fmt"
	"math"
	"math/rand"
)

const INF = math.MaxInt32 - 100

// n -> number of vertices // density -> density of the graph
func generateRandomGraph(n int, density int) [][]int {
	graph := make([][]int, n)
	for i := range graph {
		row := make([]int, n)

		for j := range row {
			if i == j {
				row[j] = 0
			} else {
				isInf := rand.Intn(density) == 0
				if isInf {
					row[j] = INF
				} else {
					row[j] = rand.Intn(10) + 1
				}
			}
		}

		graph[i] = row
	}

	return graph
}

func printGraph(graph [][]int) {
	n := len(graph)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] == INF {
				fmt.Print("INF ")
			} else {
				fmt.Printf("%d ", graph[i][j])
			}
		}
		fmt.Println()
	}
}

func printRow(row []int) {
	for i := range row {
		if row[i] == INF {
			fmt.Print("INF ")
		} else {
			fmt.Printf("%d ", row[i])
		}
	}
	fmt.Println()
}
