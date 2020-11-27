package main

import (
	"fmt"
)

func main() {
	fmt.Println("GoLang, Graph DFS and BFS implementation")
	fmt.Println("DFS : Depth First Search")
	fmt.Println("BFS : Breadth First Search")

	g := NewGraph()

	g.AddVertex("arun")
	g.AddVertex("faggy")
	g.AddVertex("mahesh")
	g.AddVertex("ariyana")

	g.AddEdge("ariyana", "arun")
	g.AddEdge("arun", "ariyana")
	g.AddEdge("faggy", "arun")
	g.AddEdge("faggy", "mahesh")
	g.AddEdge("mahesh", "faggy")
	g.AddEdge("mahesh", "ariyana")

	g.BFS("faggy")
}

func NewGraph() Graph {
	return Graph{
		adjacency: make(map[string][]string),
	}
}
type Graph struct {
	adjacency map[string][]string
}


func (g *Graph) AddVertex(vertex string) bool {
	if _, ok := g.adjacency[vertex]; ok {
		fmt.Printf("vertex %v already exists\n", vertex)
		return false
	}
	g.adjacency[vertex] = []string{}
	return true
}

func (g *Graph) AddEdge(vertex, node string) bool {
	if _, ok := g.adjacency[vertex]; !ok {
		fmt.Printf("vertex %v does not exists\n", vertex)
		return false
	}
	if ok := contains(g.adjacency[vertex], node); ok {
		fmt.Printf("node %v already exists\n", node)
		return false
	}

	g.adjacency[vertex] = append(g.adjacency[vertex], node)
	return true
}

func (g Graph) BFS(startingNode string) {
	visited := g.createVisited()
	var q []string

	visited[startingNode] = true
	q = append(q, startingNode)

	for len(q) > 0 {
		var current string
		current, q = q[0], q[1:]
		fmt.Println("visiting node", current)
		for _, node := range g.adjacency[current] {
			if !visited[node] {
				q = append(q, node)
				visited[node] = true
			}
		}
	}
}

func (g Graph) createVisited() map[string]bool {
	visited := make(map[string]bool, len(g.adjacency))
	for key := range g.adjacency {
		visited[key] = false
	}
	return visited
}


func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
} 
GoLang, Graph DFS and BFS implementation
DFS : Depth First Search
BFS : Breadth First Search
visiting node faggy
visiting node arun
visiting node mahesh
visiting node ariyana

Program exited.