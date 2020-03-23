package main

import (
	"fmt"
)

const infinite int = 100000

func main() {
	numberOfTestCases := getNumberOfTestCases()
	for i := 1; i <= numberOfTestCases; i++ {
		graph, nodeFrom, nodeTo, numberOfNodes := getTestCaseInputData()
		shortestLength := _getShortestPathLength(nodeFrom, nodeTo, graph, numberOfNodes)
		if shortestLength == infinite {
			fmt.Printf("Case #%d: no\n", i)
		} else {
			fmt.Printf("Case #%d: yes %d\n", i, shortestLength)
		}
	}
}

func _getShortestPathLength(nodeFrom int, nodeTo int, graph [][]int, numberOfNodes int) int {
	visitedNodes := make([]int, numberOfNodes, numberOfNodes)
	return getShortestPathLength(0, nodeFrom, nodeTo, visitedNodes, graph)
}

func getShortestPathLength(accumLength int, nodeFrom int, nodeTo int, visitedNodes []int, graph [][]int) int {
	var lengths []int
	visitedNodes[nodeFrom] = 1
	nextNodes := getNextNodesFor(nodeFrom, graph, visitedNodes)
	for _, nextNode := range nextNodes {
		if nextNode == nodeTo {
			lengths = append(lengths, accumLength+1)
		} else {
			lengths = append(lengths, getShortestPathLength(accumLength+1, nextNode, nodeTo, visitedNodes, graph))
		}
	}
	if len(lengths) > 0 {
		return min(lengths)
	}
	return infinite
}

func getNextNodesFor(node int, graph [][]int, visitedNodes []int) []int {
	var nextNodes []int
	row := graph[node]
	for nextNode, nodeI := range row {
		if nodeI == 1 && visitedNodes[nextNode] != 1 {
			nextNodes = append(nextNodes, nextNode)
		}
	}
	return nextNodes
}

func getTestCaseInputData() ([][]int, int, int, int) {
	numberOfNodes := getNumberOfNodes()
	graph := getGraph(numberOfNodes)
	nodeFrom, nodeTo := getFromToNodes()
	return graph, nodeFrom, nodeTo, numberOfNodes
}

func getGraph(numberOfNodes int) [][]int {
	var graph [][]int
	for i := 0; i < numberOfNodes; i++ {
		row := make([]int, numberOfNodes)
		ReadN(row, 0, numberOfNodes)
		graph = append(graph, row)
	}
	return graph
}

func getFromToNodes() (int, int) {
	fromToNodes := make([]int, 2, 2)
	ReadN(fromToNodes, 0, 2)
	return fromToNodes[0], fromToNodes[1]
}

func getNumberOfTestCases() int {
	var numberOfTestCases int
	if m, err := Scan(&numberOfTestCases); m != 1 {
		panic(err)
	}
	return numberOfTestCases
}

func getNumberOfNodes() int {
	var numberOfNodes int
	if m, err := Scan(&numberOfNodes); m != 1 {
		panic(err)
	}
	return numberOfNodes
}

func ReadN(all []int, i, n int) {
	if n == 0 {
		return
	}
	if m, err := Scan(&all[i]); m != 1 {
		panic(err)
	}
	ReadN(all, i+1, n-1)
}

func Scan(a *int) (int, error) {
	return fmt.Scan(a)
}

func min(numbers []int) int {
	m := 0
	for i, e := range numbers {
		if i==0 || e < m {
			m = e
		}
	}
	return m
}