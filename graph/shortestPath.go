package main

import (
	"fmt"
)

type ShortestPath struct {
	graph Graph
	s int
	visited []bool
	from []int
	order []int
	shortestPath []int
	stack chan int
	length []int
	lengthCount int
}
func (s *ShortestPath) initial() {
	s.stack = make(chan int, s.graph.sparseGraph.n)
	for i := 0; i < s.graph.sparseGraph.n; i ++ {
		s.visited = append(s.visited, false)
		s.from = append(s.from,-1)
		s.length = append(s.length, 0)
	}
	s.lengthCount = 1
}
func (gr *Graph) OrderRange(start int) {
	var order ShortestPath
	order.graph.sparseGraph = gr.sparseGraph
	order.s = start
	order.initial()
	order.stack <- order.s
	order.visited[order.s] = true
	order.ofs()
	fmt.Println("order:",order.order)
}

func (s *ShortestPath) ofs() {
	if len(s.stack) <= 0 {
		return
	}
	var cur int
	for i := 0; i < len(s.stack); i ++ {
		cur = <- s.stack
		s.order = append(s.order, cur)
		var arr []int
		for _, val := range s.graph.sparseGraph.g.value[cur] {
			if s.visited[val] != true {
				arr = append(arr,val)
			} else {
				continue
			}
		}
		for j := 0; j < len(arr); j ++ {
			s.stack <- arr[j]
			s.visited[arr[j]] = true
			s.length[arr[j]] = s.lengthCount
			s.from[arr[j]] = cur
		}
		s.lengthCount ++
		s.ofs()
	}
}
func (gr *Graph) ShortestPath(start int, w int) {
	var order ShortestPath
	order.graph.sparseGraph = gr.sparseGraph
	order.s = start
	order.initial()
	order.stack <- order.s
	order.visited[order.s] = true
	order.from[start] = start
	order.ofs()
	order.shortestPath = append(order.shortestPath,w)
	for i := 0; i < len(order.from); i ++ {
		if order.from[w] == start {
			order.shortestPath = append(order.shortestPath, start)
			break
		} else {
			order.shortestPath = append(order.shortestPath, order.from[w])
			w = order.from[w]
			continue
		}
	}
	j := 0
	for i := len(order.shortestPath) - 1; i >= 0; i -- {
		if i <= j {
			break
		}
		order.shortestPath[i],order.shortestPath[j] = order.shortestPath[j], order.shortestPath[i]
		j ++
	}
	fmt.Println("shortestPath:",order.shortestPath)
	sign := "->"
	for i := 0; i < len(order.shortestPath); i ++ {
		if i != len(order.shortestPath) - 1 {
			fmt.Printf("%v%v",order.shortestPath[i],sign)
		} else {
			fmt.Printf("%v",order.shortestPath[i])
		}

	}
}