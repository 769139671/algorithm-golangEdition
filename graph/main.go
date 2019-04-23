package main

import "fmt"

func main() {
	var a Graph
	a.ReadGraph("testG1.txt","sparseGraph",false)
	a.Component()
	a.IsConnected(0,7)
	a.Path(1, 4)

	var b Graph
	fmt.Println("1:")
	b.ReadGraph("testG2.txt","sparseGraph",false)
	b.OrderRange(0)
	b.ShortestPath(0,3)



}
