package main

import (
	"fmt"
)

type solution struct {
	is      integersStack
	iq      integersQueue
	graph   map[int][]int
	visited map[int]bool
}

func (s *solution) dfs(node int) {
	if !s.visited[node] {
		fmt.Println(node)
	}

	s.visited[node] = true

	for _, neighbor := range s.graph[node] { //exploring node
		if s.visited[neighbor] == true {
			continue
		}

		s.is.push(node)
		s.dfs(neighbor)
		break
	}

	for !s.is.isEmpty() {
		n, _ := s.is.pop()
		s.dfs(n)
	}
}

func (s *solution) bfs(node int) {
	fmt.Println(node)

	s.visited[node] = true
	s.iq.put(node)

	for !s.iq.isEmpty() {
		n, _ := s.iq.get()

		for _, neighbor := range s.graph[n] { // exploring node
			if s.visited[neighbor] {
				continue
			}
			s.visited[neighbor] = true // mark node as visited

			fmt.Println(neighbor)
			s.iq.put(neighbor) // put adjacent nodes in queue for future exploration
		}
	}
}

func (s *solution) resetVisited(n int) {
	for i := 0; i < n; i++ {
		s.visited[i] = false
	}
}

func (s *solution) formGraph(n int, edges [][]int) {
	for _, edge := range edges {
		s.graph[edge[0]] = append(s.graph[edge[0]], edge[1])
		s.graph[edge[1]] = append(s.graph[edge[1]], edge[0])
	}
}

func main() {

	n := 6
	edges := [][]int{{1, 0}, {1, 2}, {1, 3}, {0, 2}, {3, 4}, {4, 5}, {5, 3}}

	s := solution{integersStack{}, integersQueue{}, map[int][]int{}, map[int]bool{}}

	s.formGraph(n, edges)

	s.resetVisited(n)
	fmt.Println("DFS traversal:")
	s.dfs(0)

	s.resetVisited(n)
	fmt.Println("BFS traversal:")
	s.bfs(0)

}
