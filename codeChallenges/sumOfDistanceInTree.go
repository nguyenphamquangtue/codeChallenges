package codeChallenges

import (
	"errors"
	"fmt"
)

func ValidInputNAndEdges(n int, edges [][]int) error {
	// Validate the value of n
	if n < 1 || n > 3e4 {
		return errors.New("Invalid input N")
	}

	// Validate the length of edges
	if len(edges) != n-1 {
		return errors.New("Invalid input Edges")
	}
	return nil
}

func sumOfDistancesInTree(n int, edges [][]int) ([]int, error) {
	if err := ValidInputNAndEdges(n, edges); err != nil {
		return nil, err
	}

	// Create adjacency list to represent the tree
	adjacencyList := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adjacencyList[u] = append(adjacencyList[u], v)
		adjacencyList[v] = append(adjacencyList[v], u)
	}

	// Initialize arrays to store node count and total distances
	nodeCount := make([]int, n)
	totalDistances := make([]int, n)

	// Calculate node count and total distances from root (node 0)
	calculateNodeCountAndTotalDistances(0, -1, adjacencyList, nodeCount, totalDistances)

	// Calculate distances from root for all nodes
	calculateDistancesFromRoot(0, -1, adjacencyList, nodeCount, totalDistances, n)

	return totalDistances, nil
}

// Calculate the node count and total distances for each node using DFS
func calculateNodeCountAndTotalDistances(node, parent int, adjacencyList [][]int, nodeCount, totalDistances []int) {
	nodeCount[node] = 1
	for _, child := range adjacencyList[node] {
		if child != parent {
			calculateNodeCountAndTotalDistances(child, node, adjacencyList, nodeCount, totalDistances)
			nodeCount[node] += nodeCount[child]
			totalDistances[node] += totalDistances[child] + nodeCount[child]
		}
	}
}

// Calculate distances from root for all nodes using DFS
func calculateDistancesFromRoot(node, parent int, adjacencyList [][]int, nodeCount, totalDistances []int, n int) {
	for _, child := range adjacencyList[node] {
		if child != parent {
			totalDistances[child] = totalDistances[node] - nodeCount[child] + (n - nodeCount[child])
			calculateDistancesFromRoot(child, node, adjacencyList, nodeCount, totalDistances, n)
		}
	}
}

func RunSumOfDistanceInTree() {
	n := 4
	edges := [][]int{{0, 1}, {0, 2}, {0, 3}}

	// n := 6
	// edges := [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}
	result, err := sumOfDistancesInTree(n, edges)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
