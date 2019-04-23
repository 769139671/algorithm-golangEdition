package main

import "fmt"

type Component struct {
	G Graph
	visited []bool
	trueNum int
	ccount int
	id []int


}

func (gr *Graph) Component() {
	var com Component
	com.G.sparseGraph = gr.sparseGraph
	com.initial(gr.sparseGraph.n)
	com.component(0)
    fmt.Println("component count:",com.ccount)

}
func (gr *Graph) IsConnected(v, w int) bool {
	var com Component
	com.G.sparseGraph = gr.sparseGraph
	com.initial(gr.sparseGraph.n)
	com.component(0)

	if com.id[v] == com.id[w] {
		fmt.Println("v and w is connected")
		return true
	} else {
		fmt.Println("v and w is not connected")
		return false
	}

}
func (c *Component) initial(n int) {
	for i := 0; i < n; i ++ {
		c.visited = append(c.visited,false)
		c.id = append(c.id, -1)
	}
	c.ccount = 0
	c.trueNum = 0

}
func (c *Component) component(x int) {
	c.dfs(x)
	c.ccount ++
	if c.trueNum >=  c.G.sparseGraph.n {
		return
	} else {
		for i ,val := range c.visited {
			if val == false {
				c.component(i)
			} else {
				continue
			}
		}
	}
}

func (c *Component) dfs(x int)  {
	c.id[x] = c.ccount
	c.visited[x] = true
	c.trueNum ++
	for i := 0; i < len(c.G.sparseGraph.g.value[x]); i ++ {
		if c.visited[c.G.sparseGraph.g.value[x][i]] == false {
			c.dfs(c.G.sparseGraph.g.value[x][i])
		} else {
			continue
		}
	}
	return
}
