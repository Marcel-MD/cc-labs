package main

func floyd(graph [][]int) {
	n := len(graph)
	// Initial values of shortest
	// distances are based on shortest
	// paths considering no intermediate
	// vertex
	dist := make([][]int, n)
	for i := range dist {
		row := make([]int, n)
		for j := range row {
			row[j] = graph[i][j]
		}
		dist[i] = row
	}

	// Add all vertices one by one to
	// the set of intermediate vertices.
	for k := 0; k < n; k++ {
		// Pick all vertices as source
		// one by one
		for i := 0; i < n; i++ {
			// Pick all vertices as destination
			// for the above picked source
			for j := 0; j < n; j++ {
				// If vertex k is on the shortest
				// path from i to j, then update
				// the value of dist[i][j]
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	// printGraph(dist)
}
