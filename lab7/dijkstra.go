package main

func dijkstraAll(graph [][]int) {
	for i := range graph {
		dijkstra(graph, i)
	}
}

func dijkstra(graph [][]int, src int) {
	// initialize everything
	n := len(graph)
	dist := make([]int, n)
	spt := make([]bool, n)

	for i := 0; i < n; i++ {
		dist[i] = INF
		spt[i] = false
	}

	dist[src] = 0

	for count := 0; count < n-1; count++ {
		// Pick the minimum distance vertex
		// from the set of vertices not yet
		// processed. u is always equal to
		// src in first iteration.
		u := min(dist, spt)

		// Mark the picked vertex as processed
		spt[u] = true

		// Update dist value of the adjacent
		// vertices of the picked vertex.
		for v := 0; v < n; v++ {
			// Update dist[v] only if is not in
			// spt, there is an edge from u
			// to v, and total weight of path
			// from src to v through u is smaller
			// than current value of dist[v]
			if !spt[v] && graph[u][v] != 0 && dist[u] != INF && dist[u]+graph[u][v] < dist[v] {
				dist[v] = dist[u] + graph[u][v]
			}
		}
	}

	printRow(dist)
}

// A utility function to find the
// vertex with minimum distance
// value, from the set of vertices
// not yet included in shortest
// path tree
func min(dist []int, spt []bool) int {
	min := INF
	v := -1
	for i := range dist {
		if !spt[i] && dist[i] <= min {
			min = dist[i]
			v = i
		}
	}
	return v
}
